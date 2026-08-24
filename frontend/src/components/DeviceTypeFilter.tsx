import { createMemo, For } from "solid-js";

import { bkpHosts, filterState, setHistUpdOnFilter } from "../functions/exports";
import { filterFunc } from "../functions/filter";
import { deviceTypeFilterOptions } from "../functions/deviceTypes";

type DeviceTypeFilterProps = {
  className?: string;
  title?: string;
};

function DeviceTypeFilter(props: DeviceTypeFilterProps) {
  type FilterEvent = Event & {
    currentTarget: HTMLSelectElement;
    target: HTMLSelectElement;
  };

  const options = createMemo(() => deviceTypeFilterOptions(bkpHosts(), filterState().DeviceType));
  const active = () => filterState().DeviceType !== "";
  const selectClass = () => {
    return "form-select form-select-sm device-filter-select device-type-filter-select"
      + (props.className ? " " + props.className : "")
      + (active() ? " is-active" : "");
  };

  const handleFilter = (event: FilterEvent) => {
    filterFunc("DeviceType", event.currentTarget.value);
    setHistUpdOnFilter(true);
  };

  return (
    <select
      onChange={handleFilter}
      class={selectClass()}
      title={props.title ?? "Filter by device type"}
      value={filterState().DeviceType}
    >
      <For each={options()}>{(option) =>
        <option
          value={option.value}
          selected={filterState().DeviceType === option.value}
        >
          {option.label}
        </option>
      }</For>
    </select>
  );
}

export default DeviceTypeFilter;
