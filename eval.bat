cd %APPDATA%\JetBrains
del "PermanentDeviceId"
del "PermanentUserId"
rmdir "consentOptions" /s /q
cd %APPDATA%\JetBrains\GoLand2021.1*
rmdir "eval" /s /q
del "options\other.xml"
cd %APPDATA%\JetBrains\WebStorm2021.1*
rmdir "eval" /s /q
del "options\other.xml"
cd %APPDATA%\JetBrains\PyCharmCE2021.1*
rmdir "eval" /s /q
del "options\other.xml"
reg delete "HKEY_CURRENT_USER\Software\JavaSoft\Prefs\jetbrains"