import React from 'react'

import { BrowserRouter } from 'react-router-dom'

interface props {
  children: React.ReactNode;
}

export default (props: props) => (
  <BrowserRouter>{props.children}</BrowserRouter>
)
