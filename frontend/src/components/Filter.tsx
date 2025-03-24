import { For } from "solid-js";
import { Host, ifaces } from "../functions/exports";
import { filterFunc } from "../functions/filter";


function Filter() {

  const handleFilter = (field: keyof Host, value: any) => {
      filterFunc(field, value);
  };

  return (
    <div class="row">
      <div class="col input-group">
        <select class="form-select" title="Filter by Iface">
          <option selected disabled>Iface</option>
          <For each={ifaces()}>{(iface) =>
            <option onClick={()=>{handleFilter("Iface", iface)}}>{iface}</option>
          }</For>
        </select>
        <select class="form-select" title="Filter by Known">
          <option selected disabled>Known</option>
          <option onClick={()=>{handleFilter("Known", 1)}}>Known</option>
          <option onClick={()=>{handleFilter("Known", 0)}}>Unknown</option>
        </select>
        <select class="form-select" title="Filter by Online">
          <option selected disabled>Online</option>
          <option onClick={()=>{handleFilter("Now", 1)}}>On</option>
          <option onClick={()=>{handleFilter("Now", 0)}}>Off</option>
        </select>
        <button onClick={()=>{handleFilter("ID", 0)}} class="btn btn-outline-primary" title="Reset filter">Reset filter</button>
      </div>
    </div>
  )
}

export default Filter
