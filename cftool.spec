Name:           cftool
Version:        v1.0.0
Release:        1%{?dist}
Summary:        config help tool
Vendor:			oscar

License:		GPL
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  make

%description

%define __debug_install_post \
%{_rpmconfigdir}/find-debuginfo.sh %{?_find_debuginfo_opts} "%{_builddir}/%{?buildsubdir}"\
%{nil}

%prep
%setup -q


%build
make %{?_smp_mflags} OPTIMIZE="%{optflags}"


%install
rm -rf $RPM_BUILD_ROOT
%make_install
make install DESTDIR=%{buildroot}


%files
/usr/bin/cftool
%doc


%post


%changelog