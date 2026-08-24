import { bkpHosts, FilterState, filterState, Host, setAllHosts, setBkpHosts, sortState } from "./exports";
import { deviceTypeFilterMatches, getDeviceTypeOption, normalizeDeviceType } from "./deviceTypes";

export type HostFilterKey = keyof FilterState;

type FilterOptions = {
  ignore?: HostFilterKey[];
};

export function applyHostView() {
  const filters = filterState();
  const hosts = filterHosts(bkpHosts(), filters);

  setAllHosts(sortHosts(hosts));
}

export function updateHostInView(updatedHost: Host) {
  setBkpHosts((hosts) => hosts.map((host) => host.ID === updatedHost.ID ? { ...host, ...updatedHost } : host));
  applyHostView();
}

export function filterHosts(hosts: Host[], filters: FilterState, options: FilterOptions = {}) {
  const ignored = new Set(options.ignore ?? []);
  let filteredHosts = [...hosts];

  if (!ignored.has("Iface") && filters.Iface !== "") {
    filteredHosts = filteredHosts.filter((host) => host.Iface === filters.Iface);
  }
  if (!ignored.has("DeviceType") && filters.DeviceType !== "") {
    filteredHosts = filteredHosts.filter((host) => deviceTypeFilterMatches(host, filters.DeviceType));
  }
  if (!ignored.has("Known") && filters.Known !== "") {
    filteredHosts = filteredHosts.filter((host) => host.Known === filters.Known);
  }
  if (!ignored.has("Now") && filters.Now !== "") {
    filteredHosts = filteredHosts.filter((host) => host.Now === filters.Now);
  }
  if (!ignored.has("Search") && filters.Search !== "") {
    const search = filters.Search.toLowerCase();
    filteredHosts = filteredHosts.filter((host) => searchItem(host, search));
  }

  return filteredHosts;
}

export function hasActiveHostFilters(filters: FilterState) {
  return filters.Iface !== ""
    || filters.DeviceType !== ""
    || filters.Known !== ""
    || filters.Now !== ""
    || filters.Search.trim() !== "";
}

function sortHosts(hosts: Host[]) {
  const currentSort = sortState();

  if (!currentSort.field || !currentSort.direction) {
    return hosts;
  }

  const ascending = currentSort.direction === "ascending";
  const sortedHosts = [...hosts];

  if (currentSort.field === "IP") {
    sortedHosts.sort((a, b) => sortIP(a, b, ascending));
  } else if (currentSort.field === "DeviceType") {
    sortedHosts.sort((a, b) => byString(normalizeDeviceType(a.DeviceType), normalizeDeviceType(b.DeviceType), ascending));
  } else {
    sortedHosts.sort((a, b) => byField(a, b, currentSort.field as keyof Host, ascending));
  }

  return sortedHosts;
}

function searchItem(host: Host, search: string) {
  const name = host.Name.toLowerCase();
  const hardware = host.Hw.toLowerCase();
  const mac = host.Mac.toLowerCase();
  const iface = host.Iface.toLowerCase();
  const deviceType = getDeviceTypeOption(host.DeviceType);
  const deviceTypeValue = deviceType.value.toLowerCase();
  const deviceTypeLabel = deviceType.label.toLowerCase();

  return name.includes(search)
    || iface.includes(search)
    || host.IP.includes(search)
    || mac.includes(search)
    || hardware.includes(search)
    || deviceTypeValue.includes(search)
    || deviceTypeLabel.includes(search)
    || host.Date.includes(search);
}

function byString(a: string, b: string, ascending: boolean) {
  if (a === b) {
    return 0;
  }
  if (a > b) {
    return ascending ? 1 : -1;
  }
  return ascending ? -1 : 1;
}

function byField(a: Host, b: Host, fieldName: keyof Host, ascending: boolean) {
  if (a[fieldName] === b[fieldName]) {
    return 0;
  }
  if (a[fieldName] > b[fieldName]) {
    return ascending ? 1 : -1;
  }
  return ascending ? -1 : 1;
}

function sortIP(a: Host, b: Host, ascending: boolean) {
  const num1 = numIP(a);
  const num2 = numIP(b);
  return ascending ? num1 - num2 : num2 - num1;
}

function numIP(host: Host) {
  return Number(host.IP.split(".").map((num) => (`000${num}`).slice(-3)).join(""));
}
