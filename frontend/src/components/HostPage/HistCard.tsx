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
    <div class="card border-primary">
      <div class="card-header">
        <div class="input-group" style="width: fit-content;">
          <span class="input-group-text">Host History for</span>
          <input
            type="date"
            class="form-control"
            value={today()}
            onInput={(e) => handleDate(e.currentTarget.value)}
          />
        </div>
      </div>
      <div class="card-body">
      {_props.mac !== "" && today() !== ""
      ? <MacHistory mac={_props.mac} date={today()}></MacHistory>
      : <>Loading...</>
      }
      </div>
    </div>
  )
}

export default HistCard