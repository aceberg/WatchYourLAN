import { Host } from "../../functions/exports";
import { sortByAnyField } from "../../functions/sort";

function TableHead() {

  const handleSort = (sortBy: keyof Host) => {
    sortByAnyField(sortBy);
  };

  return (
    <thead>
      <tr>
        <th style="width: 2em;"></th>
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
  )
}

export default TableHead