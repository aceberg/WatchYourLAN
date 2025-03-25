import { For } from "solid-js";
import { Host, ifaces } from "../functions/exports";
import { filterFunc } from "../functions/filter";


function Filter() {
  type FilterEvent = Event & {
    currentTarget: HTMLSelectElement;
    target: HTMLSelectElement;
  };

  const handleFilter = (field: keyof Host, event: FilterEvent) => {
    const value = event.target ? event.target.value : 0;
    filterFunc(field, value);
  };

  return (
    <div class="row">
      <div class="col input-group">
        <select onChange={(event)=>{handleFilter("Iface", event)}} class="form-select" title="Filter by Iface">
          <option selected disabled>Iface</option>
          <For each={ifaces()}>{(iface) =>
            <option value={iface}>{iface}</option>
          }</For>
        </select>
        <select onChange={(event)=>{handleFilter("Known", event)}} class="form-select" title="Filter by Known">
          <option selected disabled>Known</option>
          <option value={1}>Known</option>
          <option value={0}>Unknown</option>
        </select>
        <select onChange={(event)=>{handleFilter("Now", event)}} class="form-select" title="Filter by Online">
          <option selected disabled>Online</option>
          <option value={1}>On</option>
          <option value={0}>Off</option>
        </select>
        <button onClick={()=>{handleFilter("ID", new Event("") as FilterEvent)}} class="btn btn-outline-primary" title="Reset filter">Reset filter</button>
      </div>
    </div>
  )
}

export default Filter
