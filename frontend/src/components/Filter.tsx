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
        <span class="input-group-text">Filter by</span>
        <select class="form-select">
          <option selected disabled>Iface</option>
          <For each={ifaces()}>{(iface) =>
            <option onClick={()=>{handleFilter("Iface", iface)}}>{iface}</option>
          }</For>
        </select>
        <select class="form-select">
          <option selected disabled>Known</option>
          <option onClick={()=>{handleFilter("Known", 1)}}>Known</option>
          <option onClick={()=>{handleFilter("Known", 0)}}>Unknown</option>
        </select>
        <select class="form-select">
          <option selected disabled>Online</option>
          <option onClick={()=>{handleFilter("Now", 1)}}>Online</option>
          <option onClick={()=>{handleFilter("Now", 0)}}>Offline</option>
        </select>
        <button onClick={()=>{handleFilter("ID", 0)}} class="btn btn-outline-primary">Reset filter</button>
      </div>
    </div>
  )
}

export default Filter
