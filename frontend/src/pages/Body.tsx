import { onMount } from "solid-js";
import { allHosts, setAllHosts } from "../functions/exports";
import { getAllHosts } from "../functions/api";
import { For } from "solid-js";
import TableRow from "../components/TableRow";


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
            <th style="width: 2em;"></th>
            </tr>
          </thead>
          <tbody>
            <For each={allHosts()}>{(host, index) =>
            <TableRow host={host} index={index}></TableRow>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default Body
