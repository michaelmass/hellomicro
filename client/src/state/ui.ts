import create from "zustand";

interface AsideState {
  minimized: boolean;
  menuOpened: string;
  sections: AsideSection[];
}

interface AsideSection {
  title?: string;
  items: AsideItem[];
}

interface AsideItem {
  text: string;
  icon: string;
  to: string;
}

export default create<AsideState>((set) => ({
  minimized: false,
  menuOpened: "",
  sections: [
    {
      items: [
        {
          text: "Dashboard",
          icon: icons.types.dashboard,
          to: routes.dashboard,
        },
        {
          text: "Status",
          icon: icons.types.status,
          to: routes.status,
        },
        {
          text: "Release",
          icon: icons.types.release,
          to: routes.release,
        },
        {
          text: "Provisions",
          icon: icons.types.provision,
          to: routes.provision,
        },
        {
          text: "Versions",
          icon: icons.types.version,
          to: routes.version,
        },
        {
          text: "Docs",
          icon: icons.types.doc,
          to: routes.doc,
        },
        {
          text: "Insight",
          icon: icons.types.insight,
          to: routes.insight,
        },
      ],
    },
  ],
}));
