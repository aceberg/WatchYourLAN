import { searchFunc } from "../functions/search";

function Search() {

  const handleSearch = (s: string) => {
      searchFunc(s);
  };

  return (
    <input onInput={e => handleSearch(e.target.value)} class="form-control" placeholder="Search" style="max-width: 10em;" title="Search"></input>
  )
}

export default Search
