import { setShow, show } from "../functions/exports";

function HistShow(_props: any) {

  const handleSaveShow = (showStr: string) => {
    localStorage.setItem(_props.name, showStr);

    const nextShow = Number(showStr);
    setShow(nextShow > 0 && !isNaN(nextShow) ? nextShow : 200);
  };

  return (
    <label class="history-samples-control">
      <span>Presence samples per device</span>
      <input
        class="form-control form-control-sm history-show-input"
        type="number"
        min="1"
        step="1"
        value={show()}
        onInput={e => handleSaveShow(e.target.value)}
        placeholder="200"
        title="Maximum number of timeline samples displayed for each device."
      ></input>
    </label>
  )
}

export default HistShow
