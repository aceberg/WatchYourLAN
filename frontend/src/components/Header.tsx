import { createSignal } from "solid-js";
import { appConfig, setAppConfig } from "../functions/exports";
import { apiGetConfig } from "../functions/api";

function Header() {

  const [themePath, setThemePath] = createSignal('');
  const [iconsPath, setIconsPath] = createSignal('');
  
  const setCurrentTheme = async () => {
    setAppConfig(await apiGetConfig());

    const theme = appConfig().Theme?appConfig().Theme:"minty";
    const color = appConfig().Color?appConfig().Color:"dark";
    
    if (appConfig().NodePath == '') {
      setThemePath("https://cdn.jsdelivr.net/npm/aceberg-bootswatch-fork@v5.3.3-2/dist/"+theme+"/bootstrap.min.css");
      setIconsPath("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.css");
    } else {
      setThemePath(appConfig().NodePath+"/node_modules/bootswatch/dist/"+theme+"/bootstrap.min.css");
      setIconsPath(appConfig().NodePath+"/node_modules/bootstrap-icons/font/bootstrap-icons.css");
    }

    document.documentElement.setAttribute("data-bs-theme", color);
    color === "dark"
      ? document.documentElement.style.setProperty('--transparent-light', '#ffffff15')
      : document.documentElement.style.setProperty('--transparent-light', '#00000015');
  }
  setCurrentTheme();

  return (
    <>
    <link rel="stylesheet" href={iconsPath()}></link> {/* icons */}
    <link rel="stylesheet" href={themePath()}></link> {/* theme */}
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
