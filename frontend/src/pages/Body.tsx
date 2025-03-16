import { onMount } from "solid-js";
import { allHosts, setAllHosts } from "../functions/exports";
import { getAllHosts } from "../functions/api";
import { For } from "solid-js";


function Body() {

  const handleSort = (sortBy: string) => {
    console.log("SORT BY", sortBy);
  };

  onMount(async () => {

    setAllHosts(await getAllHosts());
  });

  return (
    <div class="card border-primary">
      <div class="card-header">
        Header
      </div>
      <div class="card-body table-responsive">
        <table class="table table-striped">
          <thead>
            <tr>
            <th style="width: 3em;"></th>
            <th>Name <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Name"]}></i></th>
            <th>Iface <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Iface"]}></i></th>
            <th>IP <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "IP"]}></i></th>
            <th>MAC <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Mac"]}></i></th>
            <th>Hardware <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Hw"]}></i></th>
            <th>Date <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Date"]}></i></th>
            <th>Known <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Known"]}></i></th>
            <th>Online <i class="bi bi-sort-down-alt my-btn" onclick={[handleSort, "Now"]}></i></th>
            </tr>
          </thead>
          <tbody>
            <For each={allHosts()}>{(host, index) =>
            <tr>
              <td class="opacity-50">{index()+1}.</td>
              <td>{host.Name}</td>
              <td>{host.Iface}</td>
              <td><a href="http://{host.IP}" target="_blank">{host.IP}</a></td>
              <td>{host.Mac}</td>
              <td>{host.Hw}</td>
              <td>{host.Date}</td>
              <td>{host.Known}</td>
              <td>{host.Now}</td>
            </tr>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default Body
