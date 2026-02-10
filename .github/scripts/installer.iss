; TMSU Windows Installer Script (Inno Setup 6)
; Compiled by the GitHub Actions workflow — do not run standalone
; without first building tmsu.exe into the build/ directory.

#define MyAppName "TMSU"
#define MyAppVersion "0.8.0"
#define MyAppPublisher "oniony"
#define MyAppURL "https://github.com/oniony/TMSU"
#define MyAppExeName "tmsu.exe"

[Setup]
AppId={{7B2F1E4A-3C5D-4E6F-8A9B-0C1D2E3F4A5B}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}/issues
DefaultDirName={autopf}\{#MyAppName}
DefaultGroupName={#MyAppName}
DisableProgramGroupPage=yes
LicenseFile=..\..\COPYING.md
OutputDir=..\..\build
OutputBaseFilename=tmsu-windows-installer
Compression=lzma
SolidCompression=yes
WizardStyle=modern
PrivilegesRequired=lowest
ChangesEnvironment=yes

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "..\..\build\tmsu.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\..\README.md"; DestDir: "{app}"; Flags: ignoreversion
Source: "..\..\COPYING.md"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\TMSU Documentation"; Filename: "{app}\README.md"
Name: "{group}\Uninstall TMSU"; Filename: "{uninstallexe}"

[Tasks]
Name: "addtopath"; Description: "Add TMSU to your PATH environment variable"; GroupDescription: "Additional options:"

[Code]
procedure CurStepChanged(CurStep: TSetupStep);
var
  Path: string;
begin
  if (CurStep = ssPostInstall) and IsTaskSelected('addtopath') then
  begin
    RegQueryStringValue(HKCU, 'Environment', 'Path', Path);
    if Pos(ExpandConstant('{app}'), Path) = 0 then
    begin
      Path := Path + ';' + ExpandConstant('{app}');
      RegWriteStringValue(HKCU, 'Environment', 'Path', Path);
    end;
  end;
end;

procedure CurUninstallStepChanged(CurUninstallStep: TUninstallStep);
var
  Path, AppDir: string;
begin
  if CurUninstallStep = usPostUninstall then
  begin
    RegQueryStringValue(HKCU, 'Environment', 'Path', Path);
    AppDir := ExpandConstant('{app}');
    StringChangeEx(Path, ';' + AppDir, '', True);
    StringChangeEx(Path, AppDir + ';', '', True);
    StringChangeEx(Path, AppDir, '', True);
    RegWriteStringValue(HKCU, 'Environment', 'Path', Path);
  end;
end;
