import clsx from "clsx";
import { useRef, useState } from "react";
import { AiOutlineClose } from "react-icons/ai";
import { Input, InputProps } from "@/atoms/Input";
import { Space } from "@/atoms/Space";
import { List, ListEntry, TypedListColumn } from "@/atoms/List";
import { Tag } from "@/atoms/Tag";
import { Button } from "@/atoms/Button";
import { useClickoutside } from "@/hooks";
import { Theme } from "../../theme";
import "./SelectMultiple.css";

export interface SelectMultipleProps<E extends ListEntry> {
  options?: E[];
  columns?: TypedListColumn<E>[];
  listPosition?: "top" | "bottom";
  tagKey?: keyof E;
}

export const SelectMultiple = <E extends ListEntry>({
  options = [],
  columns = [],
  listPosition = "bottom",
  tagKey = "key",
}: SelectMultipleProps<E>) => {
  const [visible, setVisible] = useState(false);
  const divRef = useRef<HTMLDivElement>(null);
  useClickoutside(divRef, () => setVisible(false));
  return (
    <List data={options}>
      <Space
        className="pluto-select-multiple__container"
        ref={divRef}
        empty
        reverse={listPosition === "top"}
      >
        <List.Search
          Input={({ value, onChange }) => {
            return (
              <SelectMultipleInput
                tagKey={tagKey}
                value={value}
                focused={visible}
                onFocus={() => setVisible(true)}
                onChange={onChange}
              />
            );
          }}
        />
        <Space
          className={clsx(
            "pluto-select-multiple__list",
            `pluto-select-multiple__list--${listPosition}`,
            `pluto-select-multiple__list--${visible ? "visible" : "hidden"}`
          )}
          empty
        >
          <List.Column.Header columns={columns} />
          <List.Core.Virtual itemHeight={30}>
            {(props) => <List.Column.Item {...props} />}
          </List.Core.Virtual>
        </Space>
      </Space>
    </List>
  );
};

interface SelectMultipleInputProps<E extends ListEntry> extends InputProps {
  focused: boolean;
  onFocus: () => void;
  tagKey: keyof E;
}

const SelectMultipleInput = <E extends ListEntry>({
  value,
  onChange,
  focused,
  onFocus,
  tagKey,
}: SelectMultipleInputProps<E>) => {
  const { selected, sourceData, onSelect, clearSelected } =
    List.useContext<E>();

  const { theme } = Theme.useContext();
  return (
    <Space
      direction="horizontal"
      empty
      className="pluto-select-multiple__input__container"
      align="stretch"
      grow
    >
      <Input
        className="pluto-select-multiple__input__input"
        placeholder="Search"
        value={value}
        onChange={onChange}
        autoFocus={focused}
        onFocus={onFocus}
      />
      <Space
        direction="horizontal"
        className="pluto-select-multiple__input__tags"
        align="center"
        grow={6}
      >
        {selected
          .map((k) => sourceData.find((v) => v.key === k))
          .map((e, i) => {
            if (!e) return null;
            return (
              <Tag
                key={e.key}
                color={theme.colors.visualization.palettes.default[i]}
                onClose={() => onSelect(e.key)}
                size="small"
                variant="outlined"
              >
                {e[tagKey]}
              </Tag>
            );
          })}
      </Space>
      <Button.IconOnly
        className="pluto-select-multiple__input__tags__close"
        variant="outlined"
        onClick={clearSelected}
      >
        <AiOutlineClose aria-label="clear" />
      </Button.IconOnly>
    </Space>
  );
};
