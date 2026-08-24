import { A } from "@solidjs/router";
import { createEffect, createMemo, createSignal, For, onMount, Show } from "solid-js";

import {
  apiGetActivity,
  apiGetActivityDevices,
  apiGetActivityStats,
  type ActivityEventType,
} from "../functions/api";
import {
  activityCategoryLabel,
  activityDayLabel,
  activityDescription,
  activityDetails,
  activityDeviceIcon,
  activityDeviceTypeLabel,
  activityEventLabel,
  activityHostName,
  activityIcon,
  activityTone,
  relativeActivityTime,
} from "../functions/activity";
import { appConfig, bkpHosts, type ActivityDeviceOption, type ActivityStats, type HostEvent } from "../functions/exports";

type EventFilterKey =
  | "all"
  | "connectivity"
  | "online"
  | "offline"
  | "changes"
  | "discovered"
  | "recognition"
  | "known"
  | "unknown"
  | "device-type-changed";

type GroupByKey = "none" | "device" | "event" | "category" | "device-type" | "ip" | "iface" | "day";
type DeviceDisplayMode = "name-icon" | "name" | "icon";

type EventFilterOption = {
  key: EventFilterKey;
  label: string;
  eventTypes: ActivityEventType[];
};

type EventSummaryCard = {
  key: EventFilterKey;
  label: string;
  value: number;
  detail: string;
  icon: string;
  tone: string;
  eventTypes: ActivityEventType[];
};

type EventGroup = {
  key: string;
  label: string;
  events: HostEvent[];
};

const eventsPageSize = 100;
const customEventFilterValue = "custom";
const eventsDeviceDisplayStorageKey = "eventsDeviceDisplay";
const defaultDeviceDisplayMode: DeviceDisplayMode = "icon";
const eventTypeOrder: ActivityEventType[] = ["online", "offline", "discovered", "known", "unknown", "device-type-changed"];

const emptyStats: ActivityStats = {
  Total: 0,
  Online: 0,
  Offline: 0,
  Discovered: 0,
  Known: 0,
  Unknown: 0,
  DeviceTypeChanged: 0,
};

const eventFilterOptions: EventFilterOption[] = [
  { key: "all", label: "All events", eventTypes: [] },
  { key: "connectivity", label: "Connectivity", eventTypes: ["online", "offline"] },
  { key: "online", label: "Online", eventTypes: ["online"] },
  { key: "offline", label: "Offline", eventTypes: ["offline"] },
  { key: "changes", label: "Device changes", eventTypes: ["discovered", "known", "unknown", "device-type-changed"] },
  { key: "discovered", label: "New device detected", eventTypes: ["discovered"] },
  { key: "recognition", label: "Recognition changes", eventTypes: ["known", "unknown"] },
  { key: "known", label: "Marked known", eventTypes: ["known"] },
  { key: "unknown", label: "Marked unknown", eventTypes: ["unknown"] },
  { key: "device-type-changed", label: "Device type changed", eventTypes: ["device-type-changed"] },
];

const groupByOptions: { key: GroupByKey; label: string }[] = [
  { key: "none", label: "None" },
  { key: "device", label: "Device" },
  { key: "event", label: "Event" },
  { key: "category", label: "Category" },
  { key: "device-type", label: "Device type" },
  { key: "ip", label: "IP" },
  { key: "iface", label: "Interface" },
  { key: "day", label: "Day" },
];

const deviceDisplayOptions: { key: DeviceDisplayMode; label: string }[] = [
  { key: "name-icon", label: "Name + icon" },
  { key: "name", label: "Name only" },
  { key: "icon", label: "Icon only" },
];

function Activity() {
  const [selectedMac, setSelectedMac] = createSignal("");
  const [selectedEventTypes, setSelectedEventTypes] = createSignal<ActivityEventType[]>([]);
  const [groupBy, setGroupBy] = createSignal<GroupByKey>("none");
  const [deviceDisplayMode, setDeviceDisplayMode] = createSignal<DeviceDisplayMode>(defaultDeviceDisplayMode);
  const [events, setEvents] = createSignal<HostEvent[]>([]);
  const [stats, setStats] = createSignal<ActivityStats>(emptyStats);
  const [devices, setDevices] = createSignal<ActivityDeviceOption[]>([]);
  const [collapsedGroups, setCollapsedGroups] = createSignal<Record<string, boolean>>({});
  const [groupsCollapsedByDefault, setGroupsCollapsedByDefault] = createSignal(false);
  const [hasMore, setHasMore] = createSignal(false);
  const [loading, setLoading] = createSignal(false);
  const [error, setError] = createSignal("");
  const [statsError, setStatsError] = createSignal(false);
  let eventsRequest = 0;
  let statsRequest = 0;

  const selectedEventTypeSet = createMemo(() => new Set(selectedEventTypes()));
  const currentEventOption = () => eventFilterOptionForTypes(selectedEventTypes());
  const eventFilterSelectValue = () => currentEventOption()?.key ?? customEventFilterValue;
  const currentGroupOption = () => groupByOptions.find((option) => option.key === groupBy()) ?? groupByOptions[0];
  const hostExists = (event: HostEvent) => bkpHosts().some((host) => host.ID === event.HostID && host.Mac === event.Mac);
  const filtersActive = () => selectedMac() !== "" || selectedEventTypes().length > 0;
  const selectedDeviceLabel = () => {
    const mac = selectedMac();
    const device = devices().find((option) => option.Mac === mac);
    return device ? deviceOptionLabel(device) : mac;
  };
  const eventTypeSummary = () => selectedEventTypes().map(activityEventLabel).join(" + ");
  const eventTypeTooltipSummary = () => selectedEventTypes().map(activityEventLabel).join(", ");
  const tableStateSummary = createMemo(() => {
    const active: string[] = [];

    if (filtersActive()) {
      active.push("Filtered");
    }
    if (selectedMac() !== "") {
      active.push(selectedDeviceLabel());
    }
    if (selectedEventTypes().length > 0) {
      active.push(eventTypeSummary());
    }
    if (groupBy() !== "none") {
      active.push("Grouped by " + currentGroupOption().label);
    }

    return active.join(" · ");
  });
  const tableStateTooltip = createMemo(() => {
    const activeFilters: string[] = [];

    if (selectedMac() !== "") {
      activeFilters.push("Device: " + selectedDeviceLabel());
    }
    if (selectedEventTypes().length > 0) {
      activeFilters.push(eventTypeTooltipSummary());
    }

    const groupedBy = groupBy() !== "none" ? "Grouped by " + currentGroupOption().label : "";
    if (activeFilters.length > 0 && groupedBy) {
      return "Active filters - " + activeFilters.join(", ") + ". " + groupedBy + ".";
    }
    if (activeFilters.length > 0) {
      return "Active filters - " + activeFilters.join(", ");
    }

    return groupedBy;
  });
  const tableStateIcon = () => filtersActive() ? "bi-funnel-fill" : "bi-layers-fill";

  const retentionText = () => {
    const hours = appConfig().ConnectivityRetention || appConfig().TrimHist;
    const retention = Number.isFinite(hours) && hours > 0
      ? hours + " " + (hours === 1 ? "hour" : "hours")
      : "the configured retention window";

    return "Connectivity events retained for " + retention + ". Device changes retained while the device exists.";
  };

  const loadEvents = async (reset: boolean) => {
    const activeRequest = ++eventsRequest;
    const eventTypes = selectedEventTypes();
    const mac = selectedMac();
    const offset = reset ? 0 : events().length;

    setLoading(true);
    setError("");
    if (reset) {
      setEvents([]);
      setHasMore(false);
      setCollapsedGroups({});
    }

    try {
      const nextEvents = await apiGetActivity(eventsPageSize, {
        eventTypes: eventTypes.length === 0 ? undefined : eventTypes,
        macs: mac === "" ? undefined : [mac],
        offset,
      });
      if (activeRequest !== eventsRequest) {
        return;
      }

      setEvents((currentEvents) => reset ? nextEvents : [...currentEvents, ...nextEvents]);
      setHasMore(nextEvents.length === eventsPageSize);
    } catch {
      if (activeRequest !== eventsRequest) {
        return;
      }

      setError("Events could not be loaded");
      setHasMore(false);
    } finally {
      if (activeRequest === eventsRequest) {
        setLoading(false);
      }
    }
  };

  const loadStats = async () => {
    const activeRequest = ++statsRequest;
    const mac = selectedMac();
    setStatsError(false);

    try {
      const nextStats = await apiGetActivityStats({
        macs: mac === "" ? undefined : [mac],
      });
      if (activeRequest === statsRequest) {
        setStats(nextStats);
      }
    } catch {
      if (activeRequest === statsRequest) {
        setStats(emptyStats);
        setStatsError(true);
      }
    }
  };

  const loadDevices = async () => {
    try {
      setDevices(await apiGetActivityDevices());
    } catch {
      setDevices([]);
    }
  };

  createEffect(() => {
    selectedMac();
    selectedEventTypes();
    loadEvents(true);
  });

  createEffect(() => {
    selectedMac();
    loadStats();
  });

  onMount(() => {
    setDeviceDisplayMode(readStoredEventsDeviceDisplay());
    loadDevices();
  });

  const summaryCards = createMemo<EventSummaryCard[]>(() => {
    const currentStats = stats();
    return [
      {
        key: "all",
        label: "Total events",
        value: currentStats.Total,
        detail: statsError() ? "Counts unavailable" : "All retained events",
        icon: "bi-collection-fill",
        tone: "total",
        eventTypes: [],
      },
      {
        key: "online",
        label: "Online",
        value: currentStats.Online,
        detail: "Connectivity",
        icon: "bi-check-circle-fill",
        tone: "online",
        eventTypes: ["online"],
      },
      {
        key: "offline",
        label: "Offline",
        value: currentStats.Offline,
        detail: "Connectivity",
        icon: "bi-x-circle-fill",
        tone: "offline",
        eventTypes: ["offline"],
      },
      {
        key: "discovered",
        label: "New devices",
        value: currentStats.Discovered,
        detail: "Discovery",
        icon: "bi-plus-circle-fill",
        tone: "unknown",
        eventTypes: ["discovered"],
      },
      {
        key: "recognition",
        label: "Recognition",
        value: currentStats.Known + currentStats.Unknown,
        detail: "Known and unknown",
        icon: "bi-bookmark-check-fill",
        tone: "known",
        eventTypes: ["known", "unknown"],
      },
      {
        key: "device-type-changed",
        label: "Type changes",
        value: currentStats.DeviceTypeChanged,
        detail: "Classification",
        icon: "bi-tag-fill",
        tone: "type",
        eventTypes: ["device-type-changed"],
      },
    ];
  });

  const groupedEvents = createMemo<EventGroup[]>(() => {
    if (groupBy() === "none") {
      return [];
    }

    const groups = new Map<string, EventGroup>();
    for (const event of events()) {
      const key = groupKey(event, groupBy());
      const label = groupLabel(event, groupBy());
      const existing = groups.get(key);
      if (existing) {
        existing.events.push(event);
      } else {
        groups.set(key, { key, label, events: [event] });
      }
    }

    return [...groups.values()];
  });
  const isGroupCollapsed = (key: string) => collapsedGroups()[key] ?? groupsCollapsedByDefault();
  const hasExpandedGroups = () => groupBy() !== "none" && groupedEvents().some((group) => !isGroupCollapsed(group.key));
  const groupAllControlTitle = () => hasExpandedGroups() ? "Collapse all groups" : "Expand all groups";
  const groupAllControlIcon = () => hasExpandedGroups() ? "bi-arrows-collapse" : "bi-arrows-expand";
  const groupAllControlLabel = () => hasExpandedGroups() ? "Collapse all" : "Expand all";
  const isSummaryCardActive = (card: EventSummaryCard) => {
    if (card.key === "all") {
      return selectedEventTypes().length === 0;
    }

    const selected = selectedEventTypeSet();
    return card.eventTypes.length > 0 && card.eventTypes.every((eventType) => selected.has(eventType));
  };

  const handleReset = () => {
    setSelectedMac("");
    setSelectedEventTypes([]);
  };

  const handleSummaryClick = (event: MouseEvent, card: EventSummaryCard) => {
    if (card.key === "all") {
      setSelectedEventTypes([]);
      return;
    }

    if (event.ctrlKey || event.metaKey) {
      const selected = new Set(selectedEventTypes());
      const allCardTypesSelected = card.eventTypes.every((eventType) => selected.has(eventType));

      for (const eventType of card.eventTypes) {
        if (allCardTypesSelected) {
          selected.delete(eventType);
        } else {
          selected.add(eventType);
        }
      }

      setSelectedEventTypes(normalizeSelectedEventTypes([...selected]));
      return;
    }

    setSelectedEventTypes(
      sameEventTypes(selectedEventTypes(), card.eventTypes)
        ? []
        : normalizeSelectedEventTypes(card.eventTypes),
    );
  };

  const handleGroupToggle = (key: string) => {
    setCollapsedGroups((current) => ({
      ...current,
      [key]: !isGroupCollapsed(key),
    }));
  };

  const handleGroupAllToggle = () => {
    const nextCollapsed = hasExpandedGroups();
    const nextGroups: Record<string, boolean> = {};

    for (const group of groupedEvents()) {
      nextGroups[group.key] = nextCollapsed;
    }

    setGroupsCollapsedByDefault(nextCollapsed);
    setCollapsedGroups(nextGroups);
  };

  const handleEventFilterChange = (event: Event & { currentTarget: HTMLSelectElement }) => {
    const key = event.currentTarget.value as EventFilterKey | typeof customEventFilterValue;
    const option = eventFilterOptions.find((candidate) => candidate.key === key);

    if (option) {
      setSelectedEventTypes(normalizeSelectedEventTypes(option.eventTypes));
    }
  };

  const handleDeviceDisplayChange = (event: Event & { currentTarget: HTMLSelectElement }) => {
    const nextMode = normalizeDeviceDisplayMode(event.currentTarget.value);

    setDeviceDisplayMode(nextMode);
    persistEventsDeviceDisplay(nextMode);
  };

  return (
    <div class="activity-page">
      <header class="activity-page-header">
        <div>
          <h1 class="activity-page-title">Events</h1>
          <p class="activity-page-subtitle">Device changes and connectivity transitions</p>
        </div>
        <div class="activity-retention-note">
          <span>{retentionText()}</span>
          <A class="activity-retention-link" href="/config#data-retention">Retention settings</A>
        </div>
      </header>

      <section class="activity-summary-grid overview-grid" aria-label="Event overview">
        <For each={summaryCards()}>{(card) =>
          <button
            type="button"
            class={"overview-card overview-card-button overview-card-" + card.tone + (isSummaryCardActive(card) ? " is-active" : "")}
            aria-pressed={isSummaryCardActive(card)}
            onClick={(event) => handleSummaryClick(event, card)}
          >
            <div class="overview-card-icon" aria-hidden="true">
              <i class={"bi " + card.icon}></i>
            </div>
            <div>
              <div class="overview-card-label">{card.label}</div>
              <div class="overview-card-value">{card.value}</div>
              <div class="overview-card-detail">{card.detail}</div>
            </div>
          </button>
        }</For>
      </section>

      <section class="card wyl-panel activity-filter-panel" aria-label="Event filters">
        <div class="card-body activity-filter-grid">
          <label class="activity-filter-field">
            <span class="activity-filter-label">Device</span>
            <select
              class={"form-select form-select-sm activity-filter-select" + (selectedMac() !== "" ? " is-active" : "")}
              value={selectedMac()}
              onChange={(event) => setSelectedMac(event.currentTarget.value)}
            >
              <option value="">All devices</option>
              <For each={devices()}>{(device) =>
                <option value={device.Mac}>{deviceOptionLabel(device)}</option>
              }</For>
            </select>
          </label>
          <label class="activity-filter-field">
            <span class="activity-filter-label">Event type</span>
            <select
              class={"form-select form-select-sm activity-filter-select" + (selectedEventTypes().length > 0 ? " is-active" : "")}
              value={eventFilterSelectValue()}
              onChange={handleEventFilterChange}
            >
              <Show when={eventFilterSelectValue() === customEventFilterValue}>
                <option value={customEventFilterValue}>Multiple event types</option>
              </Show>
              <For each={eventFilterOptions}>{(option) =>
                <option value={option.key}>{option.label}</option>
              }</For>
            </select>
          </label>
          <label class="activity-filter-field">
            <span class="activity-filter-label">Group by</span>
            <select
              class={"form-select form-select-sm activity-filter-select" + (groupBy() !== "none" ? " is-active" : "")}
              value={groupBy()}
              onChange={(event) => setGroupBy(event.currentTarget.value as GroupByKey)}
            >
              <For each={groupByOptions}>{(option) =>
                <option value={option.key}>{option.label}</option>
              }</For>
            </select>
          </label>
          <label class="activity-filter-field activity-display-field">
            <span class="activity-filter-label">Device display</span>
            <select
              class="form-select form-select-sm activity-filter-select"
              value={deviceDisplayMode()}
              onChange={handleDeviceDisplayChange}
              title="Choose how the Events Device column is displayed"
            >
              <For each={deviceDisplayOptions}>{(option) =>
                <option value={option.key}>{option.label}</option>
              }</For>
            </select>
          </label>
          <Show when={groupBy() !== "none"}>
            <button
              type="button"
              class="btn btn-sm device-reset-filter activity-group-all-toggle"
              onClick={handleGroupAllToggle}
              title={groupAllControlTitle()}
              aria-label={groupAllControlTitle()}
            >
              <i class={"bi " + groupAllControlIcon()} aria-hidden="true"></i>
              <span>{groupAllControlLabel()}</span>
            </button>
          </Show>
          <Show when={filtersActive()}>
            <button type="button" class="btn btn-sm device-reset-filter activity-filter-reset" onClick={handleReset} title="Reset event filters">
              <i class="bi bi-arrow-counterclockwise" aria-hidden="true"></i>
              <span>Reset filter</span>
            </button>
          </Show>
        </div>
      </section>

      <section class="card wyl-panel activity-table-panel" aria-labelledby="events-table-title">
        <div class="card-header activity-table-header">
          <div class="activity-table-title-group">
            <div class="activity-table-title-row">
              <div id="events-table-title" class="activity-table-title">Events</div>
              <Show when={tableStateSummary()}>
                <span
                  class="activity-state-indicator"
                  title={tableStateTooltip()}
                  aria-label={tableStateTooltip()}
                >
                  <i class={"bi " + tableStateIcon()} aria-hidden="true"></i>
                  <span>{tableStateSummary()}</span>
                </span>
              </Show>
            </div>
            <div class="activity-table-subtitle">{tableSubtitle(events().length, groupBy())}</div>
          </div>
        </div>
        <div class="card-body activity-table-body">
          <div class="table-responsive">
            <table class={"table table-hover activity-table activity-table-device-display-" + deviceDisplayMode()}>
              <thead>
                <tr>
                  <th scope="col" class="activity-table-time-heading">
                    <span class="activity-heading-with-state">
                      <Show when={tableStateSummary()}>
                        <span
                          class="activity-sticky-state-indicator"
                          title={tableStateTooltip()}
                          aria-label={tableStateTooltip()}
                          role="img"
                        >
                          <i class={"bi " + tableStateIcon()} aria-hidden="true"></i>
                        </span>
                      </Show>
                      <span>Time</span>
                    </span>
                  </th>
                  <th scope="col" class="activity-table-device-heading">Device</th>
                  <th scope="col" class="activity-table-ip-heading">IP</th>
                  <th scope="col" class="activity-table-event-heading">Event</th>
                  <th scope="col" class="activity-table-iface-heading">Iface</th>
                  <th scope="col" class="activity-table-details-heading">Details</th>
                </tr>
              </thead>
              <tbody>
                <Show
                  when={groupBy() === "none"}
                  fallback={
                    <For each={groupedEvents()}>{(group) =>
                      <>
                        <tr class="activity-group-row">
                          <td colSpan={6} class="activity-group-cell" data-label="Group">
                            <button
                              type="button"
                              class="activity-group-toggle"
                              aria-expanded={!isGroupCollapsed(group.key)}
                              onClick={[handleGroupToggle, group.key]}
                            >
                              <span class="activity-group-title">
                                <i class={"bi " + (isGroupCollapsed(group.key) ? "bi-chevron-right" : "bi-chevron-down")} aria-hidden="true"></i>
                                <span>{group.label}</span>
                              </span>
                              <span class="activity-group-count">{group.events.length} loaded {group.events.length === 1 ? "event" : "events"}</span>
                            </button>
                          </td>
                        </tr>
                        <Show when={!isGroupCollapsed(group.key)}>
                          <For each={group.events}>{(event) => eventRow(event, hostExists, deviceDisplayMode)}</For>
                        </Show>
                      </>
                    }</For>
                  }
                >
                  <For each={events()}>{(event) => eventRow(event, hostExists, deviceDisplayMode)}</For>
                </Show>
              </tbody>
            </table>
          </div>
          <Show when={!loading() && events().length === 0 && error() === ""}>
            <div class="activity-empty">No events match the current filters</div>
          </Show>
          <Show when={error()}>
            <div class="activity-empty" role="status">{error()}</div>
          </Show>
          <Show when={hasMore()}>
            <div class="activity-table-actions">
              <button
                type="button"
                class="btn btn-sm wyl-button activity-load-more"
                disabled={loading()}
                onClick={() => loadEvents(false)}
              >
                <i class="bi bi-chevron-down" aria-hidden="true"></i>
                <span>{loading() ? "Loading" : "Load more"}</span>
              </button>
            </div>
          </Show>
        </div>
      </section>
    </div>
  );
}

function eventRow(event: HostEvent, hostExists: (event: HostEvent) => boolean, deviceDisplayMode: () => DeviceDisplayMode) {
  const canLinkHost = () => event.HostID > 0 && hostExists(event);
  const deviceName = () => activityHostName(event);
  const deviceTypeLabel = () => activityDeviceTypeLabel(event.DeviceType);
  const iconOnlyTitle = () => deviceName() + " \u00b7 " + deviceTypeLabel() + (canLinkHost() ? " \u00b7 open device details" : " \u00b7 Device no longer exists");
  const iconOnlyAria = () => canLinkHost()
    ? "Open " + deviceName() + " device details"
    : deviceName() + " \u00b7 " + deviceTypeLabel() + " \u00b7 Device no longer exists";
  const nameContent = () => (
    <Show
      when={canLinkHost()}
      fallback={<span class="activity-host-name">{deviceName()}</span>}
    >
      <A href={"/host/" + event.HostID} class="activity-host-link">{deviceName()}</A>
    </Show>
  );
  const deviceTypeIcon = () => (
    <span
      class="activity-host-icon"
      title={"Device type: " + deviceTypeLabel()}
      aria-label={"Device type: " + deviceTypeLabel()}
      role="img"
    >
      <i class={"bi " + activityDeviceIcon(event)} aria-hidden="true"></i>
    </span>
  );
  const iconOnlyContent = () => (
    <Show
      when={canLinkHost()}
      fallback={
        <span
          class="activity-host-icon"
          title={iconOnlyTitle()}
          aria-label={iconOnlyAria()}
          role="img"
        >
          <i class={"bi " + activityDeviceIcon(event)} aria-hidden="true"></i>
        </span>
      }
    >
      <A
        href={"/host/" + event.HostID}
        class="activity-host-icon activity-host-icon-link"
        title={iconOnlyTitle()}
        aria-label={iconOnlyAria()}
      >
        <i class={"bi " + activityDeviceIcon(event)} aria-hidden="true"></i>
      </A>
    </Show>
  );

  return (
    <tr class={"activity-table-row activity-row-" + activityTone(event.EventType)}>
      <td data-label="Time" class="activity-table-time-cell">
        <time class="activity-time" dateTime={event.Date} title={event.Date}>
          {relativeActivityTime(event.Date)}
        </time>
      </td>
      <td data-label="Device" class="activity-table-device-cell">
        <span class={"activity-table-device activity-device-display-" + deviceDisplayMode()}>
          <Show
            when={deviceDisplayMode() === "icon"}
            fallback={
              <>
                {nameContent()}
                <Show when={deviceDisplayMode() === "name-icon"}>
                  {deviceTypeIcon()}
                </Show>
              </>
            }
          >
            {iconOnlyContent()}
          </Show>
        </span>
      </td>
      <td data-label="IP" class="activity-table-muted activity-table-ip-cell">{event.IP || " "}</td>
      <td data-label="Event" class="activity-table-event-cell">
        <span class="activity-table-event">
          <span class="activity-event-icon" aria-hidden="true">
            <i class={"bi " + activityIcon(event.EventType)}></i>
          </span>
          <span>{activityDescription(event)}</span>
        </span>
      </td>
      <td data-label="Iface" class="activity-table-muted activity-table-iface-cell">{event.Iface || " "}</td>
      <td data-label="Details" class="activity-table-muted activity-table-details-cell">{activityDetails(event) || " "}</td>
    </tr>
  );
}

function tableSubtitle(count: number, groupBy: GroupByKey) {
  const base = count + " loaded " + (count === 1 ? "event" : "events");
  return groupBy === "none" ? base : base + " grouped by " + groupByOptions.find((option) => option.key === groupBy)?.label;
}

function groupKey(event: HostEvent, groupBy: GroupByKey) {
  switch (groupBy) {
    case "device":
      return "device:" + (event.Mac || event.HostID || activityHostName(event));
    case "event":
      return "event:" + event.EventType;
    case "category":
      return "category:" + activityCategoryLabel(event.EventType);
    case "device-type":
      return "device-type:" + (event.DeviceType || "not-set");
    case "ip":
      return "ip:" + (event.IP || "none");
    case "iface":
      return "iface:" + (event.Iface || "none");
    case "day":
      return "day:" + event.Date.slice(0, 10);
    case "none":
    default:
      return "all";
  }
}

function groupLabel(event: HostEvent, groupBy: GroupByKey) {
  switch (groupBy) {
    case "device":
      return activityHostName(event);
    case "event":
      return activityEventLabel(event.EventType);
    case "category":
      return activityCategoryLabel(event.EventType);
    case "device-type":
      return activityDeviceTypeLabel(event.DeviceType);
    case "ip":
      return event.IP || "No IP";
    case "iface":
      return event.Iface || "No interface";
    case "day":
      return activityDayLabel(event.Date);
    case "none":
    default:
      return "Events";
  }
}

function deviceOptionLabel(device: ActivityDeviceOption) {
  const name = device.Name.trim() || device.Mac || "Unknown device";
  return name + (device.Exists ? "" : " (deleted)");
}

function normalizeSelectedEventTypes(eventTypes: ActivityEventType[]): ActivityEventType[] {
  const selected = new Set(eventTypes);
  return eventTypeOrder.filter((eventType) => selected.has(eventType));
}

function sameEventTypes(left: ActivityEventType[], right: ActivityEventType[]) {
  const normalizedLeft = normalizeSelectedEventTypes(left);
  const normalizedRight = normalizeSelectedEventTypes(right);

  if (normalizedLeft.length !== normalizedRight.length) {
    return false;
  }

  return normalizedLeft.every((eventType, index) => eventType === normalizedRight[index]);
}

function eventFilterOptionForTypes(eventTypes: ActivityEventType[]) {
  return eventFilterOptions.find((option) => sameEventTypes(eventTypes, option.eventTypes));
}

function normalizeDeviceDisplayMode(value: unknown): DeviceDisplayMode {
  return deviceDisplayOptions.some((option) => option.key === value)
    ? value as DeviceDisplayMode
    : defaultDeviceDisplayMode;
}

function readStoredEventsDeviceDisplay(): DeviceDisplayMode {
  try {
    return normalizeDeviceDisplayMode(localStorage.getItem(eventsDeviceDisplayStorageKey));
  } catch {
    return defaultDeviceDisplayMode;
  }
}

function persistEventsDeviceDisplay(value: DeviceDisplayMode) {
  try {
    localStorage.setItem(eventsDeviceDisplayStorageKey, value);
  } catch {
    // Display preference persistence is optional; rendering should continue if storage is unavailable.
  }
}

export default Activity;
