import { createSignal, For } from "solid-js";
import { apiPortScan } from "../../functions/api";

function Ping(_props: any) {

  let stop = false;

  const [beginStr, setBegin] = createSignal("");
  const [endStr, setEnd] = createSignal("");
  const [curPort, setCurPort] = createSignal("");
  const [foundPorts, setFoundPorts] = createSignal<number[]>([]);
  const [isStopped, setIsStopped] = createSignal(false);
  const [isRunning, setIsRunning] = createSignal(false);

  const handleScan = async () => {
    stop = false;
    setIsStopped(false);
    setIsRunning(true);
    
    let begin = Number(beginStr());
    if (Number.isNaN(begin) || begin < 1 || begin > 65535) {
      begin = 1;
    }
    let end = Number(endStr());
    if (Number.isNaN(end) || end < 1 || end > 65535) {
      end = 65535;
    }

    let portOpened:boolean;
    for (let i = begin ; i <= end; i++) {

      if (stop) {
          break;
      }
      setCurPort(i.toString());
      portOpened = await apiPortScan(_props.IP, i);
      if (portOpened) {
        setFoundPorts([...foundPorts(), i]);
      }
    }
    setIsRunning(false);
  };

  const handleStop = () => {
    if (stop) {
      setBegin(curPort());
      handleScan();
    } else {
      stop = true;
      setIsStopped(true);
    }
  }

  const scanStatus = () => {
    if (isStopped()) {
      return "Paused at port: " + curPort();
    }

    if (isRunning()) {
      return "Scanning port: " + curPort();
    }

    return "Last scanned port: " + curPort();
  };

  return (
    <div class="card wyl-panel host-panel">
      <div class="card-header host-panel-header">
        <div>
          <div class="host-panel-title">Port scan</div>
          <div class="host-panel-subtitle">{_props.IP || "Waiting for host"}</div>
        </div>
      </div>
      <div class="card-body host-port-body">
        <form class="host-port-controls">
          <label class="host-port-field">
            <span>Start port</span>
            <input type="text" class="form-control form-control-sm wyl-control host-port-input" placeholder="1"
              onInput={e => setBegin(e.target.value)}></input>
          </label>
          <label class="host-port-field">
            <span>End port</span>
            <input type="text" class="form-control form-control-sm wyl-control host-port-input" placeholder="65535"
              onInput={e => setEnd(e.target.value)}></input>
          </label>
          <button type="button" onClick={handleScan} class="btn btn-sm wyl-button host-scan-button">
            <i class="bi bi-search" aria-hidden="true"></i>
            <span>Scan</span>
          </button>
        </form>
        {curPort() != ""
        ? <div class="host-scan-state">
            {isRunning() || isStopped()
            ?
            <button type="button" onClick={handleStop} class="btn btn-sm wyl-button host-stop-button">
              {isStopped() ? "Continue" : "Stop"}
            </button>
            : <></>
            }
            <div class="host-scan-status">{scanStatus()}</div>
          </div>
        : <></>
        }
        <div class="host-found-ports">
        <For each={foundPorts()}>{(port) =>
          <a class="host-port-chip" href={"http://" + _props.IP + ":" + port} target="_blank" rel="noreferrer">{port}</a>
        }</For>
        </div>
      </div>
    </div>
  )
}

export default Ping
