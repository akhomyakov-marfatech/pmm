%undefine _missing_build_ids_terminate_build

%define copying() \
%if 0%{?fedora} >= 21 || 0%{?rhel} >= 7 \
%license %{*} \
%else \
%doc %{*} \
%endif

%global repo            VictoriaMetrics
%global provider        github.com/VictoriaMetrics/%{repo}
%global commit          pmm-6401-v1.77.1

Name:           percona-victoriametrics
Version:        1.77.1
Release:        1%{?dist}
Summary:        VictoriaMetrics monitoring solution and time series database
License:        Apache-2.0
URL:            https://%{provider}
Source0:        https://%{provider}/archive/%{commit}.tar.gz


%description
%{summary}


%prep
%setup -q -n %{repo}-%{commit}


%build
export PKG_TAG=%{commit}
export BUILDINFO_TAG=%{commit}
export USER=builder

make victoria-metrics-pure
make vmalert-pure


%install
install -D -p -m 0755 ./bin/victoria-metrics-pure %{buildroot}%{_sbindir}/victoriametrics
install -D -p -m 0755 ./bin/vmalert-pure %{buildroot}%{_sbindir}/vmalert


%files
%copying LICENSE
%doc README.md
%{_sbindir}/victoriametrics
%{_sbindir}/vmalert


%changelog
* Thu May 11 2022 Michael Okoko <michael.okoko@percona.com> - 1.77.1
- upgrade victoriametrics to 1.77.1 release

* Thu Apr 14 2022 Anton Bystrov <anton.bystrov@simbirsoft.com> - 1.76.1
- upgrade victoriametrics to 1.76.1 release

* Wed Jan 20 2022 Anton Bystrov <anton.bystrov@simbirsoft.com> - 1.72.0-1
- upgrade victoriametrics to 1.72.0 release

* Thu Jun 3 2021 Vadim Yalovets <vadim.yalovets@percona.com> - 1.60.0-1
- upgrade victoriametrics to 1.60.0 release

* Wed Feb 10 2021 Vadim Yalovets <vadim.yalovets@percona.com> - 1.53.1-1
- upgrade victoriametrics to 1.53.1 release

* Wed Jan 13 2021 Vadim Yalovets <vadim.yalovets@percona.com> - 1.52.0-1
- upgrade victoriametrics to 1.52.0 release

* Thu Dec 24 2020 Vadim Yalovets <vadim.yalovets@percona.com> - 1.50.2-1
- upgrade victoriametrics to 1.50.2 release

* Tue Dec 15 2020 Nurlan Moldomurov <nurlan.moldomurov@percona.com> - 1.50.1-1
- upgrade victoriametrics to 1.50.1 release

* Tue Nov 26 2020 Nikolay Khramchikhin <nik@victoriametrics.com> - 1.48.0-1
- upgrade victoriametrics to 1.48.0 release

* Tue Nov 19 2020 Nikolay Khramchikhin <nik@victoriametrics.com> - 1.47.0-1
- upgrade victoriametrics to 1.47.0 release

* Tue Nov 10 2020 Nikolay Khramchikhin <nik@victoriametrics.com> - 1.46.0-1
- PMM-6401 upgrade victoriametrics for reading Prometheus data files

* Tue Oct 28 2020 Nikolay Khramchikhin <nik@victoriametrics.com> - 1.45.0-1
- PMM-6401 upgrade victoriametrics for reading Prometheus data files

* Tue Oct 13 2020 Aliaksandr Valialkin <valyala@victoriametrics.com> - 1.44.0-1
- PMM-6401 upgrade victoriametrics for reading Prometheus data files

* Fri Sep 4 2020 Aliaksandr Valialkin <valyala@victoriametrics.com> - 1.40.1-1
- PMM-6389 add victoriametrics and vmalert binaries
