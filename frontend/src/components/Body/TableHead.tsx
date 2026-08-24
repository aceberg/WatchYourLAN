import { For, Show } from "solid-js";
import { editNames, filterState, hasMultipleIfaces, Host, setEditNames, setSelectedIDs, SortDirection, sortState } from "../../functions/exports";
import { getHosts } from "../../functions/atstart";
import { sortByAnyField } from "../../functions/sort";
import { deviceTypeFilterLabel } from "../../functions/deviceTypes";

const headers: { label: string; field: keyof Host; className: string; title?: string; icon?: string; ariaLabel?: string }[] = [
  {
    label: "",
    field: "Known",
    className: "device-table-known",
    icon: "bi-question-circle-fill",
    title: "Known / Unknown status - click to sort",
    ariaLabel: "Known / Unknown status - click to sort",
  },
  { label: "Name", field: "Name", className: "device-table-name" },
  {
    label: "Type",
    field: "DeviceType",
    className: "device-table-type",
    title: "Device type - click to sort",
    ariaLabel: "Device type - click to sort",
  },
  { label: "IP", field: "IP", className: "device-table-ip" },
  { label: "Iface", field: "Iface", className: "device-table-iface" },
  { label: "MAC", field: "Mac", className: "device-table-mac" },
  { label: "Hardware", field: "Hw", className: "device-table-hardware" },
  {
    label: "Last Seen",
    field: "Date",
    className: "device-table-last-seen",
    title: "Last time this device was observed on the network",
  },
];

function TableHead() {
  const handleSort = (field: keyof Host) => {
    sortByAnyField(field);
  };

  const handleEditMode = async () => {
    const next = !editNames();

    if (!next) {
      await getHosts();
      setSelectedIDs([]);
    }

    setEditNames(next);
  };

  const handleKeyDown = (event: KeyboardEvent, field: keyof Host) => {
    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      handleSort(field);
    }
  };

  const ariaSort = (field: keyof Host): "ascending" | "descending" | "none" => {
    const state = sortState();
    return state.field === field && state.direction ? state.direction : "none";
  };

  const sortIcon = (direction: SortDirection | "") => {
    return direction === "ascending" ? "bi-arrow-up-short" : "bi-arrow-down-short";
  };

  const filterSummary = () => {
    const state = filterState();
    const active = [];

    if (state.Iface) {
      active.push("Iface: " + state.Iface);
    }
    if (state.DeviceType) {
      active.push("Type: " + deviceTypeFilterLabel(state.DeviceType));
    }
    if (state.Known === 1) {
      active.push("Known");
    }
    if (state.Known === 0) {
      active.push("Unknown");
    }
    if (state.Now === 1) {
      active.push("Online");
    }
    if (state.Now === 0) {
      active.push("Offline");
    }
    if (state.Search.trim()) {
      active.push('Search: "' + state.Search.trim() + '"');
    }

    return active.length > 0 ? "Active filters - " + active.join(", ") : "";
  };

  const editTitle = () => editNames() ? "Finish editing" : "Edit devices";

  return (
    <thead>
      <tr>
        <th class="device-table-index">
          <span class="device-header-content">#</span>
          <Show when={filterSummary()}>
            <span
              class="device-filter-indicator"
              title={filterSummary()}
              aria-label={filterSummary()}
              role="img"
            >
              <i class="bi bi-funnel-fill" aria-hidden="true"></i>
            </span>
          </Show>
        </th>
        <For each={headers}>{(header, index) =>
          <>
            <Show when={header.field !== "Iface" || hasMultipleIfaces()}>
            <th
              class={"sortable-th " + header.className}
              style={sortState().field === header.field ? "color: var(--wyl-link-hover);" : ''}
              aria-sort={ariaSort(header.field)}
              aria-label={header.ariaLabel}
              tabIndex={0}
              title={header.title ? header.title : "Sort by " + header.label}
              onClick={[handleSort, header.field]}
              onKeyDown={(event) => handleKeyDown(event, header.field)}
            >
              <span class="device-header-content">
                <Show when={header.icon} fallback={header.label}>
                  <i class={"bi " + header.icon} aria-hidden="true"></i>
                </Show>
              </span>
              <Show when={sortState().field === header.field}>
                <span class="device-sort-indicator" aria-hidden="true">
                  <i class={"bi " + sortIcon(sortState().direction)}></i>
                </span>
              </Show>
            </th>
            </Show>
            <Show when={index() === 0}>
              <th class="device-table-actions" title={editTitle()} aria-label={editTitle()}>
                <button
                  type="button"
                  class="device-header-action"
                  title={editTitle()}
                  aria-label={editTitle()}
                  aria-pressed={editNames()}
                  onClick={handleEditMode}
                >
                  <i class={editNames() ? "bi bi-check-lg" : "bi bi-pencil-fill"} aria-hidden="true"></i>
                </button>
              </th>
            </Show>
          </>
        }</For>
      </tr>
    </thead>
  )
}

export default TableHead
