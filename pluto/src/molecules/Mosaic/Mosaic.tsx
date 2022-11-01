import React, { useState } from "react";
import { Resize, Tab, Tabs, TabsProps } from "@/atoms";
import { MosaicLeaf } from "./types";
import { Location } from "@/util";
import "./Mosaic.css";

export interface MosaicProps extends Omit<TabsProps, "onDrop" | "tabs"> {
  onDrop: (key: number, tabKey: string, loc: Location) => void;
  onResize: (key: number, size: number) => void;
  root: MosaicLeaf;
}

export const Mosaic = (props: MosaicProps) => {
  const {
    root: { tabs, direction, first, last, key, size },
    onResize,
  } = props;

  if (tabs !== undefined) return <MosaicTabLeaf {...props} />;

  const _onResize = (sizes: number[]) => onResize(key, sizes[0]);

  if (!first || !last) {
    console.warn("Mosaic tree is malformed");
    return null;
  }

  return (
    <Resize.Multiple
      direction={direction}
      align="stretch"
      style={{
        position: "relative",
        height: "100%",
        width: "100%",
        overflow: "hidden",
      }}
      onResize={_onResize}
      initialSizes={size ? [size] : undefined}
    >
      <Mosaic {...props} root={first} />
      <Mosaic {...props} root={last} />
    </Resize.Multiple>
  );
};

const MosaicTabLeaf = ({ root: node, onDrop, ...props }: MosaicProps) => {
  const { key, tabs } = node as Omit<MosaicLeaf, "tabs"> & { tabs: Tab[] };

  const [dragMask, setDragMask] = useState<Location | null>(null);
  const [currentlyDragging, setCurrentlyDragging] = useState<string | null>(
    null
  );

  const _onDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    const validDrop =
      tabs.filter((t) => t.tabKey !== currentlyDragging).length > 0;
    if (currentlyDragging) setCurrentlyDragging(null);
    if (dragMask) setDragMask(null);
    if (!validDrop) return;
    onDrop(
      key,
      e.dataTransfer.getData("tabKey"),
      insertLocation(getDragLocationPercents(e))
    );
  };

  const onDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    const loc = insertLocation(getDragLocationPercents(e));
    // get the tab data, get a boolean value checking whether the length of the tabs
    // in node would be zero if the tab was removed
    const validDrop =
      tabs.filter((t) => t.tabKey !== currentlyDragging).length > 0;
    if (loc !== dragMask && validDrop) setDragMask(loc);
  };

  const onDragLeave = () => dragMask && setDragMask(null);

  const onDragEnter = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
  };

  return (
    <div style={{ position: "relative", height: "100%" }}>
      <Tabs
        style={{ height: "100%" }}
        tabs={tabs as Tab[]}
        {...props}
        onDrop={_onDrop}
        onDragOver={onDragOver}
        onDragLeave={onDragLeave}
        onDragEnter={onDragEnter}
        selected={node.selected}
        onTabDragStart={(e, tabEntry) => {
          e.dataTransfer.setData("tabKey", tabEntry.tabKey);
          setCurrentlyDragging(tabEntry.tabKey);
        }}
        onTabDragEnd={() => setCurrentlyDragging(null)}
      />
      {dragMask && (
        <div
          className="pluto-mosaic__drag-mask"
          style={dragMaskStyle[dragMask]}
        ></div>
      )}
    </div>
  );
};

const dragMaskStyle: Record<
  Location,
  { left: string; top: string; width: string; height: string }
> = {
  ["top"]: { left: "0%", top: "0%", width: "100%", height: "50%" },
  ["bottom"]: { left: "0%", top: "50%", width: "100%", height: "50%" },
  ["left"]: { left: "0%", top: "0%", width: "50%", height: "100%" },
  ["right"]: { left: "50%", top: "0%", width: "50%", height: "100%" },
  ["center"]: { left: "0%", top: "0%", width: "100%", height: "100%" },
};

const getDragLocationPercents = (
  e: React.DragEvent<HTMLDivElement>
): { px: number; py: number } => {
  const rect = e.currentTarget.getBoundingClientRect();
  const x = e.clientX - rect.left;
  const y = e.clientY - rect.top;
  return { px: x / rect.width, py: y / rect.height };
};

const crossHairA = (px: number): number => px;

const crossHairB = (px: number): number => 1 - px;

const insertLocation = ({ px, py }: { px: number; py: number }) => {
  if (px > 0.33 && px < 0.66 && py > 0.33 && py < 0.66) return "center";
  const [aY, bY] = [crossHairA(px), crossHairB(px)];
  if (py > aY && py > bY) return "bottom";
  if (py < aY && py < bY) return "top";
  if (py > aY && py < bY) return "left";
  if (py < aY && py > bY) return "right";
  throw new Error("[bug] - invalid insert position");
};
