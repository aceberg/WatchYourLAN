import { For, onMount, Show } from "solid-js"
import { allHosts, appConfig, hostsLoadError, setShow } from "../functions/exports"
import MacHistory from "../components/MacHistory"
import HistShow from "../components/HistShow"
import HistoryFilters from "../components/HistoryFilters"

function History() {

  onMount(() => {
    const storedShow = Number(localStorage.getItem("histShow"));
    setShow(storedShow > 0 && !isNaN(storedShow) ? storedShow : 200);
  });

  const scanIntervalHint = () => {
    const timeout = appConfig().Timeout;

    if (!Number.isFinite(timeout) || timeout <= 0) {
      return "";
    }

    return "1 block ~= " + formatInterval(timeout) + " at current scan interval.";
  };

  return (
    <>
    <Show when={hostsLoadError()}>
      <div class="data-load-warning" role="status">
        <i class="bi bi-exclamation-triangle-fill" aria-hidden="true"></i>
        <span>{hostsLoadError()}</span>
      </div>
    </Show>
    <div class="card wyl-panel history-panel">
      <div class="card-header history-panel-header">
        <div class="history-panel-title-group">
          <div class="history-panel-title">Presence</div>
          <div class="history-panel-subtitle">Sampled online/offline presence over time</div>
        </div>
        <div class="history-toolbar">
          <HistoryFilters>
            <HistShow name="histShow"></HistShow>
          </HistoryFilters>
        </div>
      </div>
      <div class="card-body history-panel-body table-responsive">
        <div class="history-legend">
          <span>Each block represents one recorded presence sample.</span>
          <Show when={scanIntervalHint()}>
            <span>{scanIntervalHint()}</span>
          </Show>
          <span class="history-legend-status"><span class="history-legend-block history-legend-online-day"></span>Online — Day</span>
          <span class="history-legend-status"><span class="history-legend-block history-legend-online-night"></span>Online — Night</span>
          <span class="history-legend-status"><span class="history-legend-block history-legend-offline"></span>Offline</span>
          <span class="history-legend-status">
            <span class="history-legend-boundary" aria-hidden="true">
              <span class="history-legend-block history-legend-neutral"></span>
              <span class="history-legend-block history-legend-neutral history-sample-hour-break"></span>
            </span>
            New hour
          </span>
          <span class="history-legend-status">
            <span class="history-legend-boundary" aria-hidden="true">
              <span class="history-legend-block history-legend-neutral"></span>
              <span class="history-legend-block history-legend-neutral history-sample-day-break"></span>
            </span>
            New day
          </span>
        </div>
        <table class="table table-hover history-table">
          <tbody>
            <For each={allHosts}>{(host, index) =>
            <tr>
              <td class="history-table-index opacity-50">{index()+1}.</td>
              <td class="history-host-cell">
                <a href={"/host/"+host.ID}>{host.Name}</a><br></br>
                <a href={"http://"+host.IP}>{host.IP}</a>
              </td>
              <td class="history-mac-cell">
                <MacHistory mac={host.Mac} date=""></MacHistory>
              </td>
            </tr>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
    </>
  )
}

function formatInterval(seconds: number) {
  if (seconds < 60) {
    return seconds + " sec";
  }

  if (seconds % 60 === 0) {
    return (seconds / 60) + " min";
  }

  return seconds + " sec";
}

export default History
