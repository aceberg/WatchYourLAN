import { createSignal, For } from "solid-js";
import { Host, ifaces, setHistUpdOnFilter } from "../functions/exports";
import { filterFunc } from "../functions/filter";


function Filter() {
  type FilterEvent = Event & {
    currentTarget: HTMLSelectElement;
    target: HTMLSelectElement;
  };

  const [selectValue, setSelectValue] = createSignal("");

  const handleFilter = (field: keyof Host, event: FilterEvent) => {
    const value = event.target ? event.target.value : 0;
    filterFunc(field, value);
    setHistUpdOnFilter(true);
  };

  const handleReset = () => {
    filterFunc("ID", 0);
    setSelectValue("something");
    setSelectValue("");
    setHistUpdOnFilter(true);
  };

  return (
    <div class="row">
      <div class="col input-group">
        <select onChange={(event)=>{handleFilter("Iface", event)}} class="form-select" title="Filter by Iface" value={selectValue()}>
          <option value="" selected disabled>Iface</option>
          <For each={ifaces()}>{(iface) =>
            <option value={iface}>{iface}</option>
          }</For>
        </select>
        <select onChange={(event)=>{handleFilter("Known", event)}} class="form-select" title="Filter by Known" value={selectValue()}>
          <option value="" selected disabled>Known</option>
          <option value={1}>Known</option>
          <option value={0}>Unknown</option>
        </select>
        <select onChange={(event)=>{handleFilter("Now", event)}} class="form-select" title="Filter by Online" value={selectValue()}>
          <option value="" selected disabled>Online</option>
          <option value={1}>On</option>
          <option value={0}>Off</option>
        </select>
        <button onClick={handleReset} class="btn btn-outline-primary" title="Reset filter">Reset filter</button>
      </div>
    </div>
  )
}

export default Filter
