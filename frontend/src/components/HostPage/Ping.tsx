import { createSignal, For } from "solid-js";
import { apiPortScan } from "../../functions/api";

function Ping(_props: any) {

  let stop = false;

  const [beginStr, setBegin] = createSignal("");
  const [endStr, setEnd] = createSignal("");
  const [curPort, setCurPort] = createSignal("");
  const [foundPorts, setFoundPorts] = createSignal<number[]>([]);

  const handleScan = async () => {
    stop = false;
    
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
  };

  const handleStop = () => {
    if (stop) {
      setBegin(curPort());
      handleScan();
    } else {
      stop = true;
    }
  }

  return (
    <div class="card border-primary">
      <div class="card-header">Port Scan</div>
      <div class="card-body">
        <form class="input-group">
          <input type="text" class="form-control" placeholder="1"
            onInput={e => setBegin(e.target.value)}></input>
          <input type="text" class="form-control" placeholder="65535"
            onInput={e => setEnd(e.target.value)}></input>
          <button type="button" onClick={handleScan} class="btn btn-primary">Scan</button>
        </form>
        {curPort() != ""
        ? <div class="d-flex justify-content-between mt-2">
            <button type="button" onClick={handleStop} class="btn btn-warning">Stop/Continue</button>
            <div>Scanning port: {curPort()}</div>
          </div>
        : <></>
        }
        <div class="mt-2">
        <For each={foundPorts()}>{(port) =>
          <a class="me-4" href={"http://" + _props.IP + ":" + port} target="_blank">{port}</a>
        }</For>
        </div>
      </div>
    </div>
  )
}

export default Ping