## [3.2.18.1] - 2018-09-07
*   chore: update sw background-uris
*   chore: update mips background-uris
*   feat(mips): set audio max volume to 150
*   feat(mips): disable right up hotzone

## [3.2.18] - 2018-08-01
*   update com.deepin.dde.launcher.override to add evince 

## [3.2.17] - 2018-08-01
*   chore: rename org.deepin.flatdeb.deepin-calendar to org.deepin.flatdeb.dde-calendar

## [3.2.16] - 2018-07-31
*   feat(mouse): Mouse accel profile option

## [3.2.15] - 2018-07-20
*   chore: workspace number defaults to 2
*   feat: com.deepin.dde.daemon add key 'trayicon' and 'x-event-monitor'
*   feat: com.deepin.dde.daemon add key 'lastore'

## [3.2.14] - 2018-06-13
*   chore(dock): update x86 dock docked-apps
*   chore(appearance): update picture-uris

<a name="3.2.13"></a>
## 3.2.13 (2018-05-30)


#### Features

* **launcher:**  update order list ([8db5855a](8db5855a))

#### Bug Fixes

* **dock:**  error type for delay-plugins-time ([d0f5588b](d0f5588b))



<a name="3.2.12"></a>
### 3.2.12 (2018-05-23)


#### Bug Fixes

* **dock:**  delay load dock plugins change to 200ms ([0ebe1b73](0ebe1b73))
* **overrides/sw:**  docked-apps no firefox ([3049a5eb](3049a5eb))

<a name="3.2.11"></a>
### 3.2.11 (2018-05-11)


#### Bug Fixes

*   correct default backgrounds path ([93613a9a](93613a9a))
*   some apps in apps-order are wrong ([65e7ef59](65e7ef59))
*   flatpak apps in apps-order are not all lowercase ([9ba106e9](9ba106e9))
*   repalce Above_Tab with other keysym string ([e464199d](e464199d))
* **sw:**  remove brightness-setter override ([5ed58ec1](5ed58ec1))

#### Features

*   add dde-welcome to zone black list ([7c059359](7c059359))
*   desktop add enable hotzone setting ([36638d65](36638d65))
*   add auto exit for deepin-notification ([23b19fbe](23b19fbe))
*   add auto exit for dde-osd ([1c1e74fc](1c1e74fc))
*   add auto-exit settings for dde-control-center ([b318a6fe](b318a6fe))
*   startdde add key iowait-enabled and memchecker-enabled ([3591218f](3591218f))
*   add calltrace ([4d1b5e0f](4d1b5e0f))
*   add debug ([a1d72dd0](a1d72dd0))
*   remove the failed grab keys from mediakey ([f55c5cf8](f55c5cf8))
*   com.deepin.dde.dameon add key 'service-trigger' ([40cb97c3](40cb97c3))
*   default-applications.terminal add key 'app-id' ([1e99107a](1e99107a))
*   dde.power add key 'adjust-brightness-enabled' ([afc1d37d](afc1d37d))
*   set swap-sched-enabled default false ([481813ba](481813ba))
*   add key swap-sched-enabled ([a54e2e48](a54e2e48))
*   add key apps-disable-scaling ([09632129](09632129))
*   remove shortcut color-picker ([19e1ab14](19e1ab14))
*   add 'auto-switch-port' in audio ([807dca57](807dca57))
*   remove deepin-movie from fullscreen-workaround-list ([407ae707](407ae707))
*   add shortcut color-picker ([1c91d4a3](1c91d4a3))
*   add shortcut deepin-system-monitor ([bcaa6105](bcaa6105))
* **dock:**  add delay-plugins-time ([bda1d76c](bda1d76c))
* **schemas:**  add key inputdevices wheel-speed ([f59912e2](f59912e2))
* **sw:**  update apps order ([efb2bd9f](efb2bd9f))
* **x86:**  add deepin flatpak apps to apps order ([ca1450d8](ca1450d8))

<a name="3.2.10"></a>
### 3.2.10 (2018-04-11)


#### Features

*   add auto exit for deepin-notification ([23b19fbe](23b19fbe))
*   add auto exit for dde-osd ([1c1e74fc](1c1e74fc))
*   add auto-exit settings for dde-control-center ([b318a6fe](b318a6fe))
*   startdde add key iowait-enabled and memchecker-enabled ([3591218f](3591218f))



## [3.2.9] - 2018-03-20
*   disable calltrace module

<a name="3.2.8"></a>
## 3.2.8 (2018-03-16)


#### Bug Fixes

*   correct default backgrounds path ([93613a9a](93613a9a))
*   some apps in apps-order are wrong ([65e7ef59](65e7ef59))
*   flatpak apps in apps-order are not all lowercase ([9ba106e9](9ba106e9))
*   repalce Above_Tab with other keysym string ([e464199d](e464199d))
* **sw:**  remove brightness-setter override ([5ed58ec1](5ed58ec1))

#### Features

*   add calltrace ([4d1b5e0f](4d1b5e0f))
*   add debug ([a1d72dd0](a1d72dd0))
*   remove the failed grab keys from mediakey ([f55c5cf8](f55c5cf8))
*   com.deepin.dde.dameon add key 'service-trigger' ([40cb97c3](40cb97c3))
*   default-applications.terminal add key 'app-id' ([1e99107a](1e99107a))
*   dde.power add key 'adjust-brightness-enabled' ([afc1d37d](afc1d37d))
*   set swap-sched-enabled default false ([481813ba](481813ba))
*   add key swap-sched-enabled ([a54e2e48](a54e2e48))
*   add key apps-disable-scaling ([09632129](09632129))
*   remove shortcut color-picker ([19e1ab14](19e1ab14))
*   add 'auto-switch-port' in audio ([807dca57](807dca57))
*   remove deepin-movie from fullscreen-workaround-list ([407ae707](407ae707))
*   add shortcut color-picker ([1c91d4a3](1c91d4a3))
*   add shortcut deepin-system-monitor ([bcaa6105](bcaa6105))
* **schemas:**  add key inputdevices wheel-speed ([f59912e2](f59912e2))
* **x86:**  add deepin flatpak apps to apps order ([ca1450d8](ca1450d8))

## [3.2.7] - 2018-03-07
+ feat: add calltrace
+ feat: add debug
+ feat: remove the failed grab keys from mediakey
+ feat: com.deepin.dde.dameon add key 'service-trigger'

## [3.2.6] - 2018-01-24
+ add 'wm' in watchdog
+ dde.power add key 'adjust-brightness-enabled'
+ default-applications.terminal add key 'app-id'

## [3.2.5] - 2018-01-10
+ add key inputdevices wheel-speed
+ sw: remove brightness-setter override
+ sw: add dde-dock preset order list
+ set swap-sched-enabled default false

## [3.2.4] - 2017-12-13
+ add key apps-disable-scaling
+ add dcc schema file
+ add key swap-sched-enabled


## [3.2.3] - 2017-11-28
+ remove shortcut color-picker


## [3.2.2] - 2017-11-16
+ correct the default backgrounds path
+ add audio port switcher

## [3.2.1] - 2017-11-09
+ change apps-order: add deepin-clone and deepin-picker and deepin flatpak apps
+ revert "fix: repalce `Above_Tab` with other keysym string"


## [3.2.0] - 2017-11-03
+ change the icon size default value
+ remove deepin-move from fullscreen workaround list
+ add shortcut for color-picker
+ add flatpak apps order list
+ add shortcut for deepin-system-monitor
+ disable touchpad palm detect


## [3.1.19] - 2017-10-25
+ add 'fprintd'

## [3.1.18] - 2017-10-23
+ replace 'Above_Tab'
+ add launcher icon ratio
