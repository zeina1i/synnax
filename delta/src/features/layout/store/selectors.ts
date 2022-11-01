import memoize from "proxy-memoize";
import { useCallback } from "react";
import { useSelector } from "react-redux";
import { LayoutStoreState } from "./slice";
import { Theme } from "@synnaxlabs/pluto";

export const useSelectLayout = (key: string) =>
  useSelector(
    useCallback(
      memoize((state: LayoutStoreState) => state.layout.layouts[key]),
      [key]
    )
  );

export const useSelectMosaic = () =>
  useSelector(
    useCallback(
      memoize((state: LayoutStoreState) => state.layout.mosaic),
      []
    )
  );

export const useSelectTheme = (): Theme => {
  const theme = useSelector(
    useCallback(
      memoize(
        (state: LayoutStoreState) => state.layout.themes[state.layout.theme]
      ),
      []
    )
  );
  return theme;
};
