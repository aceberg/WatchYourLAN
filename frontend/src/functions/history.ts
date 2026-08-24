import { apiGetHistory, apiGetHistoryByDate } from "./api";
import { Host } from "./exports";

export async function getHistoryForMac(mac: string, date: string) {
    let h:Host[] = [];
    try {
        if (date === "") {
            h = await apiGetHistory(mac);
        } else {
            h = await apiGetHistoryByDate(mac, date);
        }
    } catch {
        return [];
    }
    
    if (h != null) {
        h.sort((a:Host, b:Host) => (a.Date < b.Date ? 1 : -1));
        return h;    
    }
    return [];
}
