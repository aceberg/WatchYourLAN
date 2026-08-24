import { createEffect, createSignal, onCleanup, onMount, Show } from "solid-js";
import { A, useLocation } from "@solidjs/router";
import { appConfig, pageContext } from "../functions/exports";
import { normalizeColorMode, refreshAppConfig, setColorMode } from "../functions/theme";

function Header() {

  const [themeError, setThemeError] = createSignal(false);
  const [supportOpen, setSupportOpen] = createSignal(false);
  const location = useLocation();
  let supportButtonRef: HTMLButtonElement | undefined;
  let supportPanelRef: HTMLDivElement | undefined;
  const navItems = [
    { label: "Home", href: "/" },
    { label: "Presence", href: "/history" },
    { label: "Events", href: "/activity" },
  ];

  const currentPath = () => location.pathname.replace(/\/$/, "") || "/";
  const hostNavLabel = () => {
    const hostName = pageContext().hostName.trim();
    return hostName ? "Host · " + hostName : "Host";
  };
  const showHostContext = () => currentPath().startsWith("/host/");
  const isActivePath = (href: string) => currentPath() === (href.replace(/\/$/, "") || "/");
  const navClass = (href: string) => {
    return "nav-link wyl-nav-tab" + (isActivePath(href) ? " is-active" : "");
  };
  const settingsUtilityClass = () => "nav-link wyl-navbar-utility wyl-navbar-settings" + (isActivePath("/config") ? " is-active" : "");

  const currentColor = () => normalizeColorMode(appConfig().Color);
  const nextColor = () => currentColor() === "dark" ? "light" : "dark";
  const colorToggleLabel = () => currentColor() === "dark"
    ? "Switch to light mode"
    : "Switch to dark mode";
  const colorToggleIcon = () => currentColor() === "dark" ? "bi bi-sun-fill" : "bi bi-moon-stars-fill";

  const handleColorToggle = async () => {
    setThemeError(false);

    try {
      await setColorMode(nextColor());
    } catch (error) {
      setThemeError(true);
      console.error("Failed to save color mode", error);
    }
  };

  const handleColorToggleKeyDown = (event: KeyboardEvent) => {
    if (event.key !== "Enter" && event.key !== " ") {
      return;
    }

    event.preventDefault();
    handleColorToggle();
  };

  const closeSupport = (returnFocus = false) => {
    setSupportOpen(false);

    if (returnFocus) {
      queueMicrotask(() => supportButtonRef?.focus());
    }
  };

  const handleSupportToggle = () => {
    setSupportOpen((open) => !open);
  };

  onMount(() => {
    refreshAppConfig().catch((error) => {
      setThemeError(true);
      console.error("Failed to load application config", error);
    });

    const handlePointerDown = (event: PointerEvent) => {
      if (!supportOpen()) {
        return;
      }

      const target = event.target as Node;
      if (supportButtonRef?.contains(target) || supportPanelRef?.contains(target)) {
        return;
      }

      closeSupport();
    };

    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key !== "Escape" || !supportOpen()) {
        return;
      }

      event.preventDefault();
      closeSupport(true);
    };

    document.addEventListener("pointerdown", handlePointerDown);
    document.addEventListener("keydown", handleKeyDown);

    onCleanup(() => {
      document.removeEventListener("pointerdown", handlePointerDown);
      document.removeEventListener("keydown", handleKeyDown);
    });
  });

  createEffect(() => {
    currentPath();
    setSupportOpen(false);
  });

  return (
    <>
    <nav class="navbar navbar-expand-md navbar-dark wyl-navbar">
      <div class="container-lg">
        <a class="navbar-brand" href="/" title="LANventory" aria-label="LANventory home">
          <img src="/fs/public/lanventory-navbar.png" class="wyl-navbar-logo" alt="LANventory"/>
        </a>
        <ul class="navbar-nav wyl-nav-tabs me-auto">
          {navItems.map((item) =>
          <li class="nav-item">
            <A class={navClass(item.href)} href={item.href} title={item.label} aria-current={isActivePath(item.href) ? "page" : undefined}>{item.label}</A>
          </li>
          )}
          <Show when={showHostContext()}>
            <li class="nav-item">
              <A
                class="nav-link wyl-nav-tab wyl-nav-context is-active"
                href={currentPath()}
                title={hostNavLabel()}
                aria-current="page"
              >
                {hostNavLabel()}
              </A>
            </li>
          </Show>
        </ul>
        <ul class="navbar-nav wyl-navbar-actions">
          <li class="nav-item">
            <A
              class={settingsUtilityClass()}
              href="/config"
              title="Settings"
              aria-label="Settings"
              aria-current={isActivePath("/config") ? "page" : undefined}
            >
              <i class="bi bi-gear-fill" aria-hidden="true"></i>
            </A>
          </li>
          <li class="nav-item">
            <button
              type="button"
              class={"nav-link wyl-navbar-utility wyl-theme-toggle" + (themeError() ? " has-error" : "")}
              title={themeError() ? "Color mode could not be saved" : colorToggleLabel()}
              aria-label={themeError() ? "Color mode could not be saved" : colorToggleLabel()}
              onClick={handleColorToggle}
              onKeyDown={handleColorToggleKeyDown}
            >
              <i class={themeError() ? "bi bi-exclamation-triangle-fill" : colorToggleIcon()} aria-hidden="true"></i>
            </button>
          </li>
          <li class="nav-item wyl-support-item">
            <button
              ref={supportButtonRef}
              type="button"
              class={"nav-link wyl-navbar-utility wyl-support-toggle" + (supportOpen() ? " is-active" : "")}
              title="Support LANventory"
              aria-label="Support LANventory"
              aria-expanded={supportOpen() ? "true" : "false"}
              aria-controls="support-popover"
              onClick={handleSupportToggle}
            >
              <i class="bi bi-heart-fill" aria-hidden="true"></i>
            </button>
            <Show when={supportOpen()}>
              <div
                ref={supportPanelRef}
                id="support-popover"
                class="wyl-support-popover"
                role="dialog"
                aria-labelledby="support-popover-title"
              >
                <div id="support-popover-title" class="wyl-support-title">SUPPORT LANVENTORY</div>
                <p>LANventory is an independent fork currently in active development. If you find the project useful and want to support its continued development, you can contribute via Revolut.</p>
                <a class="wyl-support-action" href="https://revolut.me/mirgeo" target="_blank" rel="noreferrer">
                  <i class="bi bi-heart-fill" aria-hidden="true"></i>
                  <span>SUPPORT VIA REVOLUT</span>
                </a>
                <p class="wyl-support-footer">Forked from <a href="https://github.com/aceberg/WatchYourLAN" target="_blank" rel="noreferrer">WatchYourLAN by aceberg</a>.</p>
              </div>
            </Show>
          </li>
          <li class="nav-item">
            <a class="nav-link wyl-navbar-utility wyl-navbar-github" target="_blank" rel="noreferrer" href="https://github.com/godlev/WatchYourLAN2" title="LANventory on GitHub" aria-label="LANventory on GitHub"><i class="bi bi-github"></i></a>
          </li>
        </ul>
      </div>
    </nav>
    </>
  )
};

export default Header
