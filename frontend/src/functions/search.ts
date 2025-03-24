import { allHosts, bkpHosts, Host, setAllHosts } from "./exports";

export function searchFunc(s: string) {
 
  if (s != "") {

    const sl = s.toLowerCase();
    let newArray:Host[] = [];

    for (let item of allHosts) {
        
      if (searchItem(item, sl)) {
          newArray.push(item);
      }
    }

    setAllHosts(newArray);
  } else {
    setAllHosts(bkpHosts());
  }    
}

function searchItem(item:Host, sl:string) {

    const name = item.Name.toLowerCase();
    const hw = item.Hw.toLowerCase();
    const mac = item.Mac.toLowerCase();

    if ((name.includes(sl)) || (item.Iface.includes(sl)) || (item.IP.includes(sl)) || (mac.includes(sl)) || (hw.includes(sl)) || (item.Date.includes(sl))) {
        return true;
    } else {
        return false;
    }
}