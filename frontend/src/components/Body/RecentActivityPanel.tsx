import { createEffect, createMemo, createSignal, onCleanup, onMount } from "solid-js";
import { A } from "@solidjs/router";

import { apiGetActivity, type ActivityCategory } from "../../functions/api";
import { allHosts, bkpHosts, filterState, type HostEvent } from "../../functions/exports";
import { hasActiveHostFilters } from "../../functions/hostView";
import ActivityFeed from "../ActivityFeed";

const dashboardActivityLimit = 5;

type DashboardActivityPanelProps = {
  title: string;
  subtitle: string;
  events: HostEvent[];
  emptyText: string;
  hostExists: (event: HostEvent) => boolean;
};

function RecentActivityPanel() {
  const [connectivityEvents, setConnectivityEvents] = createSignal<HostEvent[]>([]);
  const [changeEvents, setChangeEvents] = createSignal<HostEvent[]>([]);
  let refreshTimer = 0;
  let requestId = 0;

  const visibleFilteredMacs = createMemo(() => {
    if (!hasActiveHostFilters(filterState())) {
      return undefined;
    }

    return allHosts.map((host) => host.Mac).filter((mac) => mac !== "");
  });

  const loadActivity = async (macs = visibleFilteredMacs()) => {
    const activeRequest = ++requestId;

    if (macs !== undefined && macs.length === 0) {
      setConnectivityEvents([]);
      setChangeEvents([]);
      return;
    }

    try {
      const [nextConnectivityEvents, nextChangeEvents] = await Promise.all([
        loadActivityCategory("connectivity", macs),
        loadActivityCategory("changes", macs),
      ]);
      if (activeRequest !== requestId) {
        return;
      }

      setConnectivityEvents(nextConnectivityEvents);
      setChangeEvents(nextChangeEvents);
    } catch {
      if (activeRequest !== requestId) {
        return;
      }

      setConnectivityEvents([]);
      setChangeEvents([]);
    }
  };

  createEffect(() => {
    loadActivity(visibleFilteredMacs());
  });

  onMount(() => {
    refreshTimer = window.setInterval(loadActivity, 60000);
  });

  onCleanup(() => {
    window.clearInterval(refreshTimer);
  });

  const hostExists = (event: HostEvent) => bkpHosts().some((host) => host.ID === event.HostID && host.Mac === event.Mac);

  return (
    <div class="activity-dashboard" aria-label="Recent events">
      <DashboardActivityPanel
        title="Connectivity"
        subtitle="Recent online/offline changes"
        events={connectivityEvents()}
        emptyText="No connectivity events recorded yet"
        hostExists={hostExists}
      ></DashboardActivityPanel>
      <DashboardActivityPanel
        title="Device changes"
        subtitle="Recent discovery and classification changes"
        events={changeEvents()}
        emptyText="No device changes recorded yet"
        hostExists={hostExists}
      ></DashboardActivityPanel>
    </div>
  );
}

async function loadActivityCategory(category: ActivityCategory, macs?: string[]) {
  return await apiGetActivity(dashboardActivityLimit, { category, macs });
}

function DashboardActivityPanel(props: DashboardActivityPanelProps) {
  return (
    <section class="card wyl-panel activity-panel activity-dashboard-panel" aria-labelledby={"dashboard-activity-" + props.title.toLowerCase().replace(/\s+/g, "-") + "-title"}>
      <div class="card-header activity-panel-header">
        <div>
          <div id={"dashboard-activity-" + props.title.toLowerCase().replace(/\s+/g, "-") + "-title"} class="activity-panel-title">{props.title}</div>
          <div class="activity-panel-subtitle">{props.subtitle}</div>
        </div>
        <A class="activity-view-all" href="/activity">
          <span>View all events</span>
          <i class="bi bi-arrow-right" aria-hidden="true"></i>
        </A>
      </div>
      <div class="card-body activity-panel-body">
        <ActivityFeed
          events={props.events}
          emptyText={props.emptyText}
          hostExists={props.hostExists}
        ></ActivityFeed>
      </div>
    </section>
  );
}

export default RecentActivityPanel;
