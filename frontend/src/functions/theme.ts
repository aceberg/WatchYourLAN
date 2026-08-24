import { apiGetConfig, apiSetConfigColor } from "./api";
import { appConfig, setAppConfig } from "./exports";
import type { Conf } from "./exports";

export type ColorMode = "dark" | "light";

const colorCacheKey = "watchyourlan-color-mode";
const defaultColor: ColorMode = "dark";
const defaultTheme = "sand";
const themeLinkID = "wyl-theme-stylesheet";

export const isColorMode = (value: unknown): value is ColorMode => value === "dark" || value === "light";

export const normalizeColorMode = (value: unknown): ColorMode => isColorMode(value) ? value : defaultColor;

export const readCachedColorMode = (): ColorMode | null => {
  try {
    const cached = localStorage.getItem(colorCacheKey);
    return isColorMode(cached) ? cached : null;
  } catch {
    return null;
  }
};

export const cacheColorMode = (color: ColorMode) => {
  try {
    localStorage.setItem(colorCacheKey, color);
  } catch {
    return;
  }
};

export const applyColorMode = (color: ColorMode) => {
  document.documentElement.setAttribute("data-bs-theme", color);
  document.documentElement.style.setProperty(
    "--transparent-light",
    color === "dark" ? "#ffffff15" : "#00000015",
  );
};

export const applyBootColorMode = () => {
  const color = readCachedColorMode() ?? defaultColor;

  applyColorMode(color);
  setAppConfig({ ...appConfig(), Color: color });
};

export const applyBaseTheme = (theme: string) => {
  const safeTheme = /^[a-z0-9-]+$/.test(theme) ? theme : defaultTheme;
  const href = "/assets/themes/" + safeTheme + "/bootstrap.min.css";
  let link = document.getElementById(themeLinkID) as HTMLLinkElement | null;

  if (!link) {
    link = document.createElement("link");
    link.id = themeLinkID;
    link.rel = "stylesheet";
    document.head.appendChild(link);
  }

  if (link.getAttribute("href") !== href) {
    link.setAttribute("href", href);
  }
};

export const applyConfigTheme = (config: Conf) => {
  const color = normalizeColorMode(config.Color);

  applyColorMode(color);
  cacheColorMode(color);
  applyBaseTheme(config.Theme || defaultTheme);
};

export const refreshAppConfig = async () => {
  const config = await apiGetConfig();
  const normalizedConfig = { ...config, Color: normalizeColorMode(config.Color) };

  setAppConfig(normalizedConfig);
  applyConfigTheme(normalizedConfig);

  return normalizedConfig;
};

export const setColorMode = async (color: ColorMode) => {
  const previousConfig = appConfig();
  const previousColor = normalizeColorMode(previousConfig.Color);
  const nextConfig = { ...previousConfig, Color: color };

  setAppConfig(nextConfig);
  applyColorMode(color);
  cacheColorMode(color);

  try {
    const persistedConfig = await apiSetConfigColor(color);
    setAppConfig(persistedConfig);
    applyConfigTheme(persistedConfig);
    return persistedConfig;
  } catch (error) {
    setAppConfig({ ...previousConfig, Color: previousColor });
    applyColorMode(previousColor);
    cacheColorMode(previousColor);
    throw error;
  }
};
