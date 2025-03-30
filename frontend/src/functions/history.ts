import { apiGetHistory } from "./api";
import { Host } from "./exports";

export async function getHistoryForMac(mac: string) {
    let h:Host[] = [];
    h = await apiGetHistory(mac);
    if (h != null) {
        h.sort((a:Host, b:Host) => (a.Date < b.Date ? 1 : -1));
        return h;    
    }
    return [];
}