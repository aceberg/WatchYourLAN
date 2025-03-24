import { apiGetAllHosts } from "./api";
import { allHosts, setAllHosts, setBkpHosts, setIfaces } from "./exports";
import { filterAtStart, filterFunc } from "./filter";
import { sortAtStart } from "./sort";

export function runAtStart() {
  getHosts();
  filterFunc("ID", 0); // reset filter

  setInterval(() => {
    getHosts();
  }, 60000); // 60000 ms = 1 minute
}

export async function getHosts() {
  const hosts = await apiGetAllHosts();

  if (hosts !== null && hosts.length > 0) {
    setAllHosts(hosts);
    setBkpHosts(hosts);

    listIfaces();
    sortAtStart();
    filterAtStart();
  }
}

function listIfaces() {

  let ifaces:string[] = [];

  for (let host of allHosts) {
    if (!ifaces.includes(host.Iface)) {
      ifaces.push(host.Iface);
    }
  }

  setIfaces(ifaces);
}