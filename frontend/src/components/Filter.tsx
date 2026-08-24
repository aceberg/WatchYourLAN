import { For, Show } from "solid-js";
import { filterState, hasMultipleIfaces, Host, ifaces, setHistUpdOnFilter } from "../functions/exports";
import { filterFunc, resetFilters } from "../functions/filter";
import { hasActiveHostFilters } from "../functions/hostView";
import DeviceTypeFilter from "./DeviceTypeFilter";


function Filter() {
  type FilterEvent = Event & {
    currentTarget: HTMLSelectElement;
    target: HTMLSelectElement;
  };

  const handleFilter = (field: keyof Host, event: FilterEvent) => {
    const value = event.currentTarget.value;
    filterFunc(field, value);
    setHistUpdOnFilter(true);
  };

  const handleReset = () => {
    resetFilters();
    setHistUpdOnFilter(true);
  };

  const hasActiveFilter = () => {
    return hasActiveHostFilters(filterState());
  };

  return (
    <div class="device-filter-group">
      <Show when={hasMultipleIfaces()}>
        <select
          onChange={(event)=>{handleFilter("Iface", event)}}
          class={"form-select form-select-sm device-filter-select" + (filterState().Iface !== "" ? " is-active" : "")}
          title="Filter by Iface"
          value={filterState().Iface}
        >
          <option value="">Iface</option>
          <For each={ifaces()}>{(iface) =>
            <option value={iface}>{iface}</option>
          }</For>
        </select>
      </Show>
        <DeviceTypeFilter title="Filter by device type"></DeviceTypeFilter>
        <Show when={hasActiveFilter()}>
          <button onClick={handleReset} class="btn btn-sm device-reset-filter" title="Reset filter">
            <i class="bi bi-arrow-counterclockwise" aria-hidden="true"></i>
            <span>Reset filter</span>
          </button>
        </Show>
    </div>
  )
}

export default Filter
