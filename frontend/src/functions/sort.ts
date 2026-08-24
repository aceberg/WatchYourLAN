import { Host, setSortState, sortState, SortDirection } from "./exports";
import { applyHostView } from "./hostView";

const sortableFields = ["Name", "DeviceType", "Iface", "IP", "Mac", "Hw", "Date", "Known"];

export function sortAtStart() {
  const field = normalizeField(localStorage.getItem("sortField"));
  const direction = readStoredDirection();

  if (!field || !direction) {
    clearSortState();
    return;
  }

  setSortState({ field, direction });
}

export function sortByAnyField(field: keyof Host) {
  const current = sortState();

  if (current.field !== field || current.direction === "") {
    setSort(field, "ascending");
  } else if (current.direction === "ascending") {
    setSort(field, "descending");
  } else {
    clearSortState();
  }

  applyHostView();
}

function setSort(field: keyof Host, direction: SortDirection) {
  localStorage.setItem("sortField", field);
  localStorage.setItem("sortDown", (direction === "ascending").toString());
  setSortState({ field, direction });
}

function clearSortState() {
  localStorage.removeItem("sortField");
  localStorage.removeItem("sortDown");
  setSortState({ field: "", direction: "" });
}

function normalizeField(field: string | null): keyof Host | "" {
  return field && sortableFields.includes(field) ? field as keyof Host : "";
}

function readStoredDirection(): SortDirection | "" {
  const stored = localStorage.getItem("sortDown");
  if (stored === "true") return "ascending";
  if (stored === "false") return "descending";
  return "";
}
