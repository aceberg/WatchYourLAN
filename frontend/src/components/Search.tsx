import { filterState } from "../functions/exports";
import { searchFunc } from "../functions/search";

type SearchProps = {
  className?: string;
  onSearch?: () => void;
  placeholder?: string;
  title?: string;
};

function Search(props: SearchProps) {
  type SearchEvent = InputEvent & {
    currentTarget: HTMLInputElement;
    target: HTMLInputElement;
  };

  const label = () => props.title ?? props.placeholder ?? "Search";
  const inputClass = () => "form-control form-control-sm device-search" + (props.className ? " " + props.className : "");

  const handleSearch = (s: string) => {
    searchFunc(s);
    props.onSearch?.();
  };

  return (
    <input
      onInput={(event: SearchEvent) => handleSearch(event.currentTarget.value)}
      value={filterState().Search}
      class={inputClass()}
      placeholder={props.placeholder ?? "Search"}
      title={label()}
      aria-label={label()}
    ></input>
  )
}

export default Search
