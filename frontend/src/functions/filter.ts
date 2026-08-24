import { emptyFilterState, filterState, FilterState, Host, setFilterState } from "./exports";
import { applyHostView } from "./hostView";
import { normalizeDeviceTypeFilter } from "./deviceTypes";

type FilterField = "Iface" | "DeviceType" | "Known" | "Now";

export function filterAtStart() {
  setFilterState((current) => ({
    ...current,
    ...readStoredFilters(),
  }));
}

export function filterFunc(field: keyof Host, value: any) {
  if (field === "ID") {
    resetFilters();
    return;
  }

  if (field !== "Iface" && field !== "DeviceType" && field !== "Known" && field !== "Now") {
    return;
  }

  setHostFilter(field, value);
}

export function toggleHostFilter(field: "Known" | "Now", value: number) {
  const current = filterState();
  const nextValue = current[field] === value ? "" : value;

  updateFilterState({
    ...current,
    [field]: nextValue,
  });
}

export function resetFilters() {
  updateFilterState(emptyFilterState);
}

export function setSearchFilter(value: string) {
  updateFilterState({
    ...filterState(),
    Search: value,
  });
}

function setHostFilter(field: FilterField, value: any) {
  const normalized = field === "Iface"
    ? normalizeIface(value)
    : field === "DeviceType"
      ? normalizeDeviceTypeFilter(value)
      : normalizeBinary(value);

  updateFilterState({
    ...filterState(),
    [field]: normalized,
  });
}

function updateFilterState(nextState: FilterState) {
  setFilterState(nextState);
  persistFilters(nextState);
  applyHostView();
}

function readStoredFilters(): Omit<FilterState, "Search"> {
  const storedFilters = {
    Iface: normalizeIface(localStorage.getItem("filterIface")),
    DeviceType: normalizeDeviceTypeFilter(localStorage.getItem("filterDeviceType")),
    Known: normalizeBinary(localStorage.getItem("filterKnown")),
    Now: normalizeBinary(localStorage.getItem("filterNow")),
  };

  if (storedFilters.Iface !== ""
    || storedFilters.DeviceType !== ""
    || storedFilters.Known !== ""
    || storedFilters.Now !== "") {
    return storedFilters;
  }

  const legacyField = localStorage.getItem("filterField") as FilterField | null;
  const legacyValue = localStorage.getItem("filterValue");

  if (legacyField === "Iface") {
    return { ...storedFilters, Iface: normalizeIface(legacyValue) };
  }
  if (legacyField === "DeviceType") {
    return { ...storedFilters, DeviceType: normalizeDeviceTypeFilter(legacyValue) };
  }
  if (legacyField === "Known" || legacyField === "Now") {
    return { ...storedFilters, [legacyField]: normalizeBinary(legacyValue) };
  }

  return storedFilters;
}

function persistFilters(state: FilterState) {
  persistValue("filterIface", state.Iface);
  persistValue("filterDeviceType", state.DeviceType);
  persistValue("filterKnown", state.Known);
  persistValue("filterNow", state.Now);
  localStorage.removeItem("filterField");
  localStorage.removeItem("filterValue");
}

function persistValue(key: string, value: string | number) {
  if (value === "") {
    localStorage.removeItem(key);
  } else {
    localStorage.setItem(key, value.toString());
  }
}

function normalizeIface(value: any) {
  return typeof value === "string" ? value : "";
}

function normalizeBinary(value: any): number | "" {
  if (value === 1 || value === "1") return 1;
  if (value === 0 || value === "0") return 0;
  return "";
}
