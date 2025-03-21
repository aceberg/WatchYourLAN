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

async function getHosts() {
  const hosts = await apiGetAllHosts();
  setAllHosts(hosts);
  setBkpHosts(hosts);

  sortAtStart();
  filterAtStart();
  listIfaces();
  // console.log("UPD");
}

function listIfaces() {

  let ifaces:string[] = [];

  for (let host of allHosts()) {
    if (!ifaces.includes(host.Iface)) {
      ifaces.push(host.Iface);
    }
  }

  setIfaces(ifaces);
}