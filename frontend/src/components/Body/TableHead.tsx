import { createSignal, For } from "solid-js";
import { Host } from "../../functions/exports";
import { sortByAnyField } from "../../functions/sort";

function TableHead() {

  const [sortField, setSortField] = createSignal<string>('');
  
  const showSort = () => {
    let field = localStorage.getItem("sortField") as string;
    field === "Mac" ? field = "MAC" : '';
    field === "Hw" ? field = "Hardware" : '';
    field === "Now" ? field = "On" : '';
    setSortField(field);
  };
  showSort();

  const handleSort = (sortBy: string) => {
    setSortField(sortBy);
    sortBy === "MAC" ? sortBy = "Mac" : '';
    sortBy === "Hardware" ? sortBy = "Hw" : '';
    sortBy === "On" ? sortBy = "Now" : '';
    sortByAnyField(sortBy as keyof Host);
  };

  return (
    <thead>
      <tr>
        <th style="width: 2em;"></th>
        <For each={["Name", "Iface", "IP", "MAC", "Hardware", "Date", "Known", "On"]}>{(key) =>
          <th 
            style={key === sortField() ? "color: var(--bs-primary);" : ''}
          >{key} <i
            class="bi bi-sort-down-alt my-btn"
            onClick={[handleSort, key]}
            title={"Sort by " + key}
          ></i></th>
        }</For>
        <th style="width: 2em;" title="Edit"><i class="bi bi-pencil-fill"></i></th>
      </tr>
    </thead>
  )
}

export default TableHead