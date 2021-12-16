export const host = process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:3000' : `https://${location.hostname}`

export const icons = {
  types: {
    echo: '',
  },
  actions: {
    unknown: '',
    create: '/images/svg/icons/Navigation/Plus.svg',
    edit: '/images/svg/icons/Communication/Write.svg',
    delete: '/images/svg/icons/General/Trash.svg',
    remove: '/images/svg/icons/Communication/Forward.svg',
    add: '/images/svg/icons/Code/Plus.svg',
    copy: '/images/svg/icons/General/Duplicate.svg',
    sync: '/images/svg/icons/Media/Repeat.svg',
    search: '/images/svg/icons/General/Search.svg',
    params: '/images/svg/icons/Shopping/Settings.svg',
    settings: '/images/svg/icons/General/Settings-2.svg',
    review: '/images/svg/icons/General/Binocular.svg',
    cancel: '/images/svg/icons/Navigation/Close.svg',
    confirm: '/images/svg/icons/Navigation/Check.svg',
    lock: '/images/svg/icons/General/Lock.svg',
    unlock: '/images/svg/icons/General/Unlock.svg',
    log: '/images/svg/icons/Text/Align-left.svg',
    logout: '/images/svg/icons/Navigation/Sign-out.svg',
    filter: '/images/svg/icons/Text/Filter.svg',
  },
  aside: {
    logo: '/images/logos/logo.svg',
    name: '/images/logos/name.svg',
    arrow: '/images/svg/icons/Navigation/Angle-double-left.svg',
  },
  arrows: {
    up: '/images/svg/icons/Navigation/Up-2.svg',
    down: '/images/svg/icons/Navigation/Down-2.svg',
    left: '/images/svg/icons/Navigation/Left-2.svg',
    right: '/images/svg/icons/Navigation/Right-2.svg',
  },
  sort: {
    asc: '/images/svg/icons/Navigation/Up-2.svg',
    desc: '/images/svg/icons/Navigation/Down-2.svg',
    default: '/images/svg/icons/Shopping/Sort1.svg',
  },
  error: {
    unauthorized: '/images/svg/error/unauthorized.svg',
    notfound: '/images/svg/error/notfound.svg',
    internal: '/images/svg/error/internal.svg',
  },
}

export enum BaseRoute {
  Echo = 'echo',
}

export const routes: { [route in BaseRoute]: string } = {
  [BaseRoute.Echo]: '/',
}

export const routesIcons: { [route in BaseRoute]: string } = {
  [BaseRoute.Echo]: icons.types.echo,
}
