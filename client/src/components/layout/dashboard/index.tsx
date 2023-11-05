import React from "react";

import { Scrollbar } from "components/elements";
import { Loadbar } from "components/preset";

import { Header } from "components/layout/header";
import { SubHeader } from "components/layout/subheader";
import { SideMenu } from "components/layout/sidemenu";

import Aside from "./aside";

interface props {
  children: React.ReactNode;
}

export const Dashboard = (props: props) => (
  <div className="flex bg-gray-100 w-screen min-h-screen overflow-hidden">
    <div
      id="app"
      className="absolute overflow-hidden w-screen h-screen pointer-events-none"
    />
    <Loadbar />
    <Aside />
    <div className="flex flex-col flex-1 h-screen dark:text-white dark:bg-dark overflow-hidden">
      <Header />
      <SubHeader />
      <Scrollbar className="-mt-12" options={{ suppressScrollX: true }}>
        <div className="py-5 xl:px-20 px-5 flex-1">
          {props.children}
        </div>
      </Scrollbar>
    </div>
    <SideMenu />
  </div>
);
