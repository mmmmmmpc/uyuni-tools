#
# spec file for package uyunictl
#
# Copyright (c) 2023 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           uyunictl
Version:        0.1
Release:        0%{?dist}
Summary:	Client tool for Uyuni and SUSE Manager
License:        GPLv2
URL:            https://github.com/uyuni-project/uyuni-tools
Source:         https://github.com/uyuni-project/uyuni-tools
BuildRequires:  go

%description
The uyunictl is an API Client application written in Go that connects to the Uyuni and/or SUSE Manager server API and perform day to day user tasks such as gathering system information.

%prep
# Nothing to do here

%build
# Build the Go application
%{gobuild} %{name}

%install
# Create the installation directory
install -d %{buildroot}/%{_bindir}

# Install the binary
install -m 0755 %{name} %{buildroot}/%{_bindir}/%{name}

%files
%{_bindir}/%{name}
%license COPYING
%doc ChangeLog README

%post
%postun

%changelog
* %{version}-%{release}
- Initial package
