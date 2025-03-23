import { For } from "solid-js";
import { Host } from "../../functions/exports";
import { sortByAnyField } from "../../functions/sort";

function TableHead() {

  const handleSort = (sortBy: string) => {
    sortBy === "MAC" ? sortBy = "Mac" : '';
    sortBy === "Hardware" ? sortBy = "Hw" : '';
    sortBy === "Online" ? sortBy = "Now" : '';
    sortByAnyField(sortBy as keyof Host);
  };

  return (
    <thead>
      <tr>
        <th style="width: 2em;"></th>
        <For each={["Name", "Iface", "IP", "MAC", "Hardware", "Date", "Known", "Online"]}>{(key) =>
          <th>{key} <i class="bi bi-sort-down-alt my-btn" 
                      onClick={[handleSort, key]}
                      title={"Sort by " + key}
                    ></i></th>
        }</For>
        <th style="width: 2em;"></th>
      </tr>
    </thead>
  )
}

export default TableHead