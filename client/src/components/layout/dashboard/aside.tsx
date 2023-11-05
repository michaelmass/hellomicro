import React from "react";

import { useRecoilState } from "recoil";
import { asideState } from "state/ui";
import { Aside } from "components/layout/aside";

export default () => {
  const [state, setState] = useRecoilState(asideState);

  return (
    <Aside
      minimized={state.minimized}
      setMinimized={(minimized: boolean) => setState({ ...state, minimized })}
    >
      {state.sections.map((section, i) => (
        <React.Fragment key={i}>
          {section.title && (
            <Aside.Section minimized={state.minimized} text={section.title} />
          )}
          {section.items.map((item, j) => (
            <Aside.Item
              key={j}
              minimized={state.minimized}
              menuOpened={state.menuOpened}
              text={item.text}
              icon={item.icon}
              to={item.to}
            />
          ))}
        </React.Fragment>
      ))}
    </Aside>
  );
};
