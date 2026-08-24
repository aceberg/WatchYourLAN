import { createSignal, onMount } from "solid-js";
import { setShow } from "../../functions/exports";
import MacHistory from "../MacHistory"

function HistCard(_props: any) {

  const [today, setToday] = createSignal('');

  onMount(() => {
    setShow(15000);
    setToday(new Date().toLocaleDateString("en-CA"));
  });

  const handleDate = (date: string) => {
    setToday("");
    setToday(date);
  };

  return (
    <div class="card wyl-panel host-history-panel">
      <div class="card-header host-history-header">
        <div>
          <div class="host-panel-title">Presence history</div>
          <div class="host-panel-subtitle">{_props.mac || "Waiting for host"}</div>
        </div>
        <label class="host-history-date-control">
          <span>Presence date</span>
          <input
            type="date"
            class="form-control form-control-sm wyl-control host-date-input"
            value={today()}
            onInput={(e) => handleDate(e.currentTarget.value)}
          />
        </label>
      </div>
      <div class="card-body host-history-body">
      {_props.mac !== "" && today() !== ""
      ? <MacHistory mac={_props.mac} date={today()}></MacHistory>
      : <span class="host-loading">Loading...</span>
      }
      </div>
    </div>
  )
}

export default HistCard
