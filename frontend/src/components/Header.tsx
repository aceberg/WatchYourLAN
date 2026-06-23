import { onMount } from "solid-js";
import { appConfig, setAppConfig } from "../functions/exports";
import { apiGetConfig } from "../functions/api";

function Header() {
  
  const setCurrentTheme = async () => {
    setAppConfig(await apiGetConfig());

    const theme = appConfig().Theme?appConfig().Theme:"sand";
    const color = appConfig().Color?appConfig().Color:"dark";
    let themePath = "/fs/public/vendor/aceberg-bootswatch-fork/dist/"+theme+"/bootstrap.min.css";
    let iconsPath = "/fs/public/vendor/bootstrap-icons/font/bootstrap-icons.min.css";
    
    if (appConfig().NodePath != '') {
      themePath = appConfig().NodePath+"/node_modules/bootswatch/dist/"+theme+"/bootstrap.min.css";
      iconsPath = appConfig().NodePath+"/node_modules/bootstrap-icons/font/bootstrap-icons.css";
    }

    const themeLink = document.getElementById("watchyourlan-theme-css") as HTMLLinkElement | null;
    const iconsLink = document.getElementById("watchyourlan-icons-css") as HTMLLinkElement | null;
    themeLink?.setAttribute("href", themePath);
    iconsLink?.setAttribute("href", iconsPath);

    document.documentElement.setAttribute("data-bs-theme", color);
    color === "dark"
      ? document.documentElement.style.setProperty('--transparent-light', '#ffffff15')
      : document.documentElement.style.setProperty('--transparent-light', '#00000015');
  }

  onMount(() => {
    setCurrentTheme();
  });

  return (
    <>
    <nav class="navbar navbar-expand-md navbar-dark bg-primary">
      <div class="container-lg">
        <a class="navbar-brand" href="/">
          <img src="/fs/public/favicon.png" style="width: 2em"/>
        </a>
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <a class="nav-link active" href="/" title="Home">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" href="/config/" title="Config">Config</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" href="/history/" title="History">History</a>
          </li>
        </ul>
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link active fs-3 ms-md-2" target="_blank" href="https://github.com/aceberg/WatchYourLAN" title="Github"><i class="bi bi-github"></i></a>
          </li>
        </ul>
      </div>
    </nav>
    </>
  )
};

export default Header
