import { setShow, show } from "../functions/exports";

function HistShow(_props: any) {

  const handleSaveShow = (showStr: string) => {
    localStorage.setItem(_props.name, showStr);

    setShow(+showStr);
    show() == 0 ? setShow(200) : '';
  };

  return (
    <input class="form-control" onInput={e => handleSaveShow(e.target.value)} placeholder="Show elements" title="Nomber of elements to show" style="max-width: 10em;"></input>
  )
}

export default HistShow
