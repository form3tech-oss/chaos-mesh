# Changelog

All notable changes to this project will be documented in this file. See [commit-and-tag-version](https://github.com/absolute-version/commit-and-tag-version) for commit guidelines.

## [2.0.1](https://github.com/form3tech-oss/chaos-mesh/compare/v2.0.0...v2.0.1) (2023-06-30)

## 2.0.0 (2023-06-30)


### ⚠ BREAKING CHANGES

* build binaries locally with `local/` prefix targets in `Makefile` (#4004)
* support PhysicalMachine in UI (#3624)
* UI monorepo (#2590)

### Features

* Add `StatusCheck` CRD ([#2954](https://github.com/form3tech-oss/chaos-mesh/issues/2954)) ([7c114d6](https://github.com/form3tech-oss/chaos-mesh/commit/7c114d676d3ad9ace3ded0b9dec0ad0466c63567))
* add archive report page ([#783](https://github.com/form3tech-oss/chaos-mesh/issues/783)) ([35ef5e2](https://github.com/form3tech-oss/chaos-mesh/commit/35ef5e2fbb5f7252af2ca8a4f38a917fbb194ee7))
* add events page ([#628](https://github.com/form3tech-oss/chaos-mesh/issues/628)) ([a1e7fa4](https://github.com/form3tech-oss/chaos-mesh/commit/a1e7fa49047270f272644cc20a1ae709d4d5562e))
* add global search for events, experiments and archives. ([#1034](https://github.com/form3tech-oss/chaos-mesh/issues/1034)) ([db7ce35](https://github.com/form3tech-oss/chaos-mesh/commit/db7ce357bfacefc7ff264ed53467805a8ee72f85))
* add hostNetwork option for controller-manager and update doc accordingly ([#874](https://github.com/form3tech-oss/chaos-mesh/issues/874)) ([2cc3393](https://github.com/form3tech-oss/chaos-mesh/commit/2cc3393c92b6d849e784f91357881243724cb92e))
* add junit report xml output ([#2170](https://github.com/form3tech-oss/chaos-mesh/issues/2170)) ([9d8c35b](https://github.com/form3tech-oss/chaos-mesh/commit/9d8c35b73db0f588e199d6d881b0464f8c8f4aa2))
* add PhysicalMachineChaos UI ([#2417](https://github.com/form3tech-oss/chaos-mesh/issues/2417)) ([3d6bea1](https://github.com/form3tech-oss/chaos-mesh/commit/3d6bea175619df5546d2d38e9c8cf9905cbde77e))
* add pprof register and config option for chaos dashboard. ([#1071](https://github.com/form3tech-oss/chaos-mesh/issues/1071)) ([711d81c](https://github.com/form3tech-oss/chaos-mesh/commit/711d81c2a0a1bfccab81ce345ab3f4cfd104e5bc))
* also append workflow name for controlled chaos object ([#1900](https://github.com/form3tech-oss/chaos-mesh/issues/1900)) ([820b693](https://github.com/form3tech-oss/chaos-mesh/commit/820b69301d654882f50c7643d13f5971a3f2edb4))
* apis for workflow endtime and status ([#1828](https://github.com/form3tech-oss/chaos-mesh/issues/1828)) ([9efe08c](https://github.com/form3tech-oss/chaos-mesh/commit/9efe08c2ca4586f4a8ea7d6aeb3ce8efc363eec8))
* apply multi type of chaos nodes with code generation ([#1621](https://github.com/form3tech-oss/chaos-mesh/issues/1621)) ([936d764](https://github.com/form3tech-oss/chaos-mesh/commit/936d7646c72631b54e561a1c34b5aa3ddbe3cef8))
* base implementation of parallel node ([#1827](https://github.com/form3tech-oss/chaos-mesh/issues/1827)) ([9088aea](https://github.com/form3tech-oss/chaos-mesh/commit/9088aeaa1199391e007b25f8769cca61a00eb07f))
* build binaries locally with `local/` prefix targets in `Makefile` ([#4004](https://github.com/form3tech-oss/chaos-mesh/issues/4004)) ([e2f2292](https://github.com/form3tech-oss/chaos-mesh/commit/e2f2292d8d4e14da7da9b94939262f9e3c9bb513))
* **chaos dashboard:** add deleting status ([#2708](https://github.com/form3tech-oss/chaos-mesh/issues/2708)) ([0528b02](https://github.com/form3tech-oss/chaos-mesh/commit/0528b027d9d4a728a3f87d60cf860c4b46f853df))
* chaosctl for print debug info ([#1074](https://github.com/form3tech-oss/chaos-mesh/issues/1074)) ([d9c7291](https://github.com/form3tech-oss/chaos-mesh/commit/d9c7291bcfce274dcb8b3fa2d40ec9fd0740999e))
* compose the name of task and node ([#1853](https://github.com/form3tech-oss/chaos-mesh/issues/1853)) ([c5c1bac](https://github.com/form3tech-oss/chaos-mesh/commit/c5c1bac9c9f603f449e986e9b3026ec56bf6dad9))
* dashboard API of task node ([#2062](https://github.com/form3tech-oss/chaos-mesh/issues/2062)) ([ab694ec](https://github.com/form3tech-oss/chaos-mesh/commit/ab694ecbb961cd4bc881d6155845a9f9e87f1b64))
* enable chaos-controller-manager leader election ([#2291](https://github.com/form3tech-oss/chaos-mesh/issues/2291)) ([9375ba9](https://github.com/form3tech-oss/chaos-mesh/commit/9375ba9111ce2f020e3f4d86a18be1542e6d295c))
* enable mTLS between chaos-controller-manager and chaosd ([#2580](https://github.com/form3tech-oss/chaos-mesh/issues/2580)) ([15ba206](https://github.com/form3tech-oss/chaos-mesh/commit/15ba2062036c6995b6f70ec284ef773f46ee531b))
* enriching events for workflow ([#2083](https://github.com/form3tech-oss/chaos-mesh/issues/2083)) ([ae407ad](https://github.com/form3tech-oss/chaos-mesh/commit/ae407ad59ccb217e5ea044db0e24689b99ef1076))
* experiment creation ([#577](https://github.com/form3tech-oss/chaos-mesh/issues/577)) ([3be4666](https://github.com/form3tech-oss/chaos-mesh/commit/3be4666e31f27ddd3c9e0f8736cb1047e5194cb0))
* experiment detail ([#593](https://github.com/form3tech-oss/chaos-mesh/issues/593)) ([e5698da](https://github.com/form3tech-oss/chaos-mesh/commit/e5698da63e911e3aa3cf15980fd987cd4e7070a6))
* i18n support and dark mode ([#898](https://github.com/form3tech-oss/chaos-mesh/issues/898)) ([df4cf9e](https://github.com/form3tech-oss/chaos-mesh/commit/df4cf9ec6cc18a7f9dd66576cf44165e31512da0))
* make selector support expression selectors. ([#1277](https://github.com/form3tech-oss/chaos-mesh/issues/1277)) ([e97ff52](https://github.com/form3tech-oss/chaos-mesh/commit/e97ff52401192fb0b8361dbdff8b52a39d5c5543))
* **metrics:** add archived objects metrics for chaos-dashboard ([#2568](https://github.com/form3tech-oss/chaos-mesh/issues/2568)) ([bdd4588](https://github.com/form3tech-oss/chaos-mesh/commit/bdd4588313a96cf6f0ced5b035bd80e9e5e46b3b))
* **metrics:** add bpm controlled processes metrics ([#2497](https://github.com/form3tech-oss/chaos-mesh/issues/2497)) ([ddfa197](https://github.com/form3tech-oss/chaos-mesh/commit/ddfa1972fe5c4e094c4e3ef79775600a85632135))
* **metrics:** add emitted event counter metric ([#2435](https://github.com/form3tech-oss/chaos-mesh/issues/2435)) ([bf27d3e](https://github.com/form3tech-oss/chaos-mesh/commit/bf27d3eb6e4111db989fdcf55a28a92a297b26dc))
* **metrics:** add gRPC client request duration histogram metric ([#2458](https://github.com/form3tech-oss/chaos-mesh/issues/2458)) ([b85bf5a](https://github.com/form3tech-oss/chaos-mesh/commit/b85bf5a90e4f940ad5c44d352083c983820b1456))
* **metrics:** add iptables, ipset and tc metrics ([#2540](https://github.com/form3tech-oss/chaos-mesh/issues/2540)) ([dd26ace](https://github.com/form3tech-oss/chaos-mesh/commit/dd26acef735ef71fdb788deb8ef56e6af3d7561f))
* **metrics:** add schedule and workflow metrics ([#2402](https://github.com/form3tech-oss/chaos-mesh/issues/2402)) ([c0f3dec](https://github.com/form3tech-oss/chaos-mesh/commit/c0f3decff120d81e10e19aab12d53dfd9276102b))
* **metrics:** gRPC and HTTP request duration histogram metrics ([#2543](https://github.com/form3tech-oss/chaos-mesh/issues/2543)) ([0d1d27f](https://github.com/form3tech-oss/chaos-mesh/commit/0d1d27f9187f5da1309f252556fae48935a2fd23))
* more settings in dashboard ([#1449](https://github.com/form3tech-oss/chaos-mesh/issues/1449)) ([1118b9e](https://github.com/form3tech-oss/chaos-mesh/commit/1118b9e7c5372c946cad377029bb4d8d9408b486))
* new ci for upload coverage data to codecov ([#2679](https://github.com/form3tech-oss/chaos-mesh/issues/2679)) ([a5744c6](https://github.com/form3tech-oss/chaos-mesh/commit/a5744c6186570114e69fc300a706651006403218))
* next `New Workflow` in UI ([#3185](https://github.com/form3tech-oss/chaos-mesh/issues/3185)) ([26ec509](https://github.com/form3tech-oss/chaos-mesh/commit/26ec5093990ded0e7d137d5b12972e544df35d9b))
* OpenAPI to TypeScript API Client and Form Data ([#2770](https://github.com/form3tech-oss/chaos-mesh/issues/2770)) ([48e27e6](https://github.com/form3tech-oss/chaos-mesh/commit/48e27e6983d36165ddf29dc19928adb6767e1325))
* remove helm v2 tiller and use helm 3.5.3 as default ([#1575](https://github.com/form3tech-oss/chaos-mesh/issues/1575)) ([0e27b1e](https://github.com/form3tech-oss/chaos-mesh/commit/0e27b1ee82e043d4f972d8dc4678be10f73aded0))
* setup OWNERS for chaos-mesh ([#4039](https://github.com/form3tech-oss/chaos-mesh/issues/4039)) ([29335e8](https://github.com/form3tech-oss/chaos-mesh/commit/29335e8e65550b8df7828721326cd6e5f232a64a))
* show event and yaml component for  workflow ([#2275](https://github.com/form3tech-oss/chaos-mesh/issues/2275)) ([a4c56f5](https://github.com/form3tech-oss/chaos-mesh/commit/a4c56f5745e74e1b54e552321b41952a6130e6b3))
* **stresschaos:** support cgroup v2 for docker and cri-o ([#3698](https://github.com/form3tech-oss/chaos-mesh/issues/3698)) ([e26351d](https://github.com/form3tech-oss/chaos-mesh/commit/e26351d6495dcb0803fd59a8d7b5ee2f196084c5))
* support awschaos in UI ([#1682](https://github.com/form3tech-oss/chaos-mesh/issues/1682)) ([6d1d58e](https://github.com/form3tech-oss/chaos-mesh/commit/6d1d58e1e89b5d835b27edca62b9b071c9cd0b3e))
* support batch delete in UI ([#1723](https://github.com/form3tech-oss/chaos-mesh/issues/1723)) ([8ad8a6f](https://github.com/form3tech-oss/chaos-mesh/commit/8ad8a6ff66abbceb0707e89e2169b08e8b7692fa))
* support cgroup v2 for linux stress experiments ([#2928](https://github.com/form3tech-oss/chaos-mesh/issues/2928)) ([57832f7](https://github.com/form3tech-oss/chaos-mesh/commit/57832f791d75592104a54313e46ffa976082c4c6)), closes [#2937](https://github.com/form3tech-oss/chaos-mesh/issues/2937) [#2935](https://github.com/form3tech-oss/chaos-mesh/issues/2935) [#2915](https://github.com/form3tech-oss/chaos-mesh/issues/2915) [#2770](https://github.com/form3tech-oss/chaos-mesh/issues/2770) [#2911](https://github.com/form3tech-oss/chaos-mesh/issues/2911) [#2921](https://github.com/form3tech-oss/chaos-mesh/issues/2921) [#2947](https://github.com/form3tech-oss/chaos-mesh/issues/2947) [#2824](https://github.com/form3tech-oss/chaos-mesh/issues/2824) [#2919](https://github.com/form3tech-oss/chaos-mesh/issues/2919) [#2918](https://github.com/form3tech-oss/chaos-mesh/issues/2918) [#2902](https://github.com/form3tech-oss/chaos-mesh/issues/2902) [#2948](https://github.com/form3tech-oss/chaos-mesh/issues/2948)
* support DNSChaos in dashboard ([#1238](https://github.com/form3tech-oss/chaos-mesh/issues/1238)) ([2ab3b19](https://github.com/form3tech-oss/chaos-mesh/commit/2ab3b19ea5054f90d7cfca99254a95f46f7d9294))
* support gcp authentication on UI ([#2363](https://github.com/form3tech-oss/chaos-mesh/issues/2363)) ([61359c8](https://github.com/form3tech-oss/chaos-mesh/commit/61359c8994dbf02981d9a948e149a53124405ed2))
* support GcpChaos in dashboard ([#1686](https://github.com/form3tech-oss/chaos-mesh/issues/1686)) ([f1c817e](https://github.com/form3tech-oss/chaos-mesh/commit/f1c817e07ea339d41ded06055998c9a1e9f135e9))
* support new 2.0 features on the UI ([#1977](https://github.com/form3tech-oss/chaos-mesh/issues/1977)) ([c644613](https://github.com/form3tech-oss/chaos-mesh/commit/c64461332ab9d789d8c15d3a395e3edf5e51871f))
* support PhysicalMachine in UI ([#3624](https://github.com/form3tech-oss/chaos-mesh/issues/3624)) ([5749dc0](https://github.com/form3tech-oss/chaos-mesh/commit/5749dc01775afa2a61783cae17d2586b715438d7))
* switch views between k8s and hosts nodes ([#3830](https://github.com/form3tech-oss/chaos-mesh/issues/3830)) ([1a807ff](https://github.com/form3tech-oss/chaos-mesh/commit/1a807ff9d3f2ebceb98a39b5cb0218e3e62b905b))
* token generator ([#1507](https://github.com/form3tech-oss/chaos-mesh/issues/1507)) ([7798d83](https://github.com/form3tech-oss/chaos-mesh/commit/7798d83e7afc98cd6459d3f8c0523587939f48fe))
* UI monorepo ([#2590](https://github.com/form3tech-oss/chaos-mesh/issues/2590)) ([ae52a08](https://github.com/form3tech-oss/chaos-mesh/commit/ae52a08c8203a9555359de65deb11718dac954ce))
* **ui:** add storybook for testing ([#2994](https://github.com/form3tech-oss/chaos-mesh/issues/2994)) ([0842a2d](https://github.com/form3tech-oss/chaos-mesh/commit/0842a2d4ff01ecc982450b663deb8c618958bb37))
* **ui:** allow importing external workflows and copying flow nodes in next generation `New Workflow` ([#3368](https://github.com/form3tech-oss/chaos-mesh/issues/3368)) ([e0e0d54](https://github.com/form3tech-oss/chaos-mesh/commit/e0e0d54f646362d8df81826b4b3b3c723302779a))
* **ui:** support `Suspend` in next generation `New Workflow` ([#3254](https://github.com/form3tech-oss/chaos-mesh/issues/3254)) ([43d5781](https://github.com/form3tech-oss/chaos-mesh/commit/43d57818db6483d2b8fe1713b4d3b0a376c6f348))
* update API requests with OpenAPI generated client ([#2926](https://github.com/form3tech-oss/chaos-mesh/issues/2926)) ([046961d](https://github.com/form3tech-oss/chaos-mesh/commit/046961de0af65f20a61c569888952dea4256bd36)), closes [#3386](https://github.com/form3tech-oss/chaos-mesh/issues/3386)
* update the start time of workflow ([#1862](https://github.com/form3tech-oss/chaos-mesh/issues/1862)) ([aeda04a](https://github.com/form3tech-oss/chaos-mesh/commit/aeda04aec991bbda014acf5b588f9b5de25a74c3))
* use YAML format as experiment description ([#1029](https://github.com/form3tech-oss/chaos-mesh/issues/1029)) ([4c55adb](https://github.com/form3tech-oss/chaos-mesh/commit/4c55adb081889bd7e4c6addcffeb77f8477351eb))
* validation webhook for workflow ([#2028](https://github.com/form3tech-oss/chaos-mesh/issues/2028)) ([6ab935d](https://github.com/form3tech-oss/chaos-mesh/commit/6ab935d213e24c2b09ea2dcc0de8b04418106354))
* workflows UI ([#1870](https://github.com/form3tech-oss/chaos-mesh/issues/1870)) ([d65a973](https://github.com/form3tech-oss/chaos-mesh/commit/d65a973122e7d150d204eb7520391409c891d50e))


### Bug Fixes

* **#2155:** 404 not found ([#2156](https://github.com/form3tech-oss/chaos-mesh/issues/2156)) ([1eb4025](https://github.com/form3tech-oss/chaos-mesh/commit/1eb4025c68580e16045bba4273b07d50768feffc)), closes [#2155](https://github.com/form3tech-oss/chaos-mesh/issues/2155)
* add `envFollowKubernetesPattern` to handle k8s-like format env in helm templates ([#2955](https://github.com/form3tech-oss/chaos-mesh/issues/2955)) ([d0b0236](https://github.com/form3tech-oss/chaos-mesh/commit/d0b0236962d30aa897e651f4d9741d383e3ce5d7))
* add max validation to workers in stress chaos ([#2010](https://github.com/form3tech-oss/chaos-mesh/issues/2010)) ([407aef2](https://github.com/form3tech-oss/chaos-mesh/commit/407aef2d390ea02b0265f972f3222af004bc6f28))
* add some annotation to pull request template ([#2276](https://github.com/form3tech-oss/chaos-mesh/issues/2276)) ([d75b7fd](https://github.com/form3tech-oss/chaos-mesh/commit/d75b7fdaef1b2a89cc293a9895fd9c22dca51005))
* add the missing `context.Context` param in bpm/build_darwin ([#2996](https://github.com/form3tech-oss/chaos-mesh/issues/2996)) ([d215648](https://github.com/form3tech-oss/chaos-mesh/commit/d21564823806e8c019cad91be693eeea40dd53df))
* also enter pid namespace for 'cat /proc/mounts'  ([#1307](https://github.com/form3tech-oss/chaos-mesh/issues/1307)) ([643db9f](https://github.com/form3tech-oss/chaos-mesh/commit/643db9f1aaf2f6048ac14dcba8760e6f9a0b3c86))
* also set NextRecover with cron expression ([#1588](https://github.com/form3tech-oss/chaos-mesh/issues/1588)) ([dc0bd3c](https://github.com/form3tech-oss/chaos-mesh/commit/dc0bd3c3ff5b76307e88ade72918897e095b21dd))
* always use latest for build-env and dev-env ([#2582](https://github.com/form3tech-oss/chaos-mesh/issues/2582)) ([4c4d792](https://github.com/form3tech-oss/chaos-mesh/commit/4c4d7929c977e9e533837c1a5275a8d6e0d041b5))
* api compatibility of dashboard ([#752](https://github.com/form3tech-oss/chaos-mesh/issues/752)) ([f96be09](https://github.com/form3tech-oss/chaos-mesh/commit/f96be09e59da0cccc8f76a74603d0f22861b173b))
* back to the old version of kind ([#1634](https://github.com/form3tech-oss/chaos-mesh/issues/1634)) ([ac48aba](https://github.com/form3tech-oss/chaos-mesh/commit/ac48abaad6e246aad5de5c86f5a5ce5a23967794))
* broken links in docs ([#741](https://github.com/form3tech-oss/chaos-mesh/issues/741)) ([4ea62b3](https://github.com/form3tech-oss/chaos-mesh/commit/4ea62b3dab60c673ff9d4c12a9c723696e4831a0))
* bundle ui assets when build dashboard ([#2456](https://github.com/form3tech-oss/chaos-mesh/issues/2456)) ([edc1945](https://github.com/form3tech-oss/chaos-mesh/commit/edc1945c09aa465aa3e01bf95336dfad80bb4462))
* can not delete last token ([#1873](https://github.com/form3tech-oss/chaos-mesh/issues/1873)) ([307eb10](https://github.com/form3tech-oss/chaos-mesh/commit/307eb10363a0ccc72a1f28474e1de0f94b1d3815))
* can not skip e2e test ([#3749](https://github.com/form3tech-oss/chaos-mesh/issues/3749)) ([9685160](https://github.com/form3tech-oss/chaos-mesh/commit/9685160ce2e3081781a0d658f147427b4c83a264))
* can't load schedule archives ([#2524](https://github.com/form3tech-oss/chaos-mesh/issues/2524)) ([4b1bee0](https://github.com/form3tech-oss/chaos-mesh/commit/4b1bee0583becf91badf7c4ed12df146cf351766))
* can't set to immediate job ([#1233](https://github.com/form3tech-oss/chaos-mesh/issues/1233)) ([9954eab](https://github.com/form3tech-oss/chaos-mesh/commit/9954eab0c72f1a51971058a8693819372334d938))
* **chaos daemon:** logging kvs as pair ([#3716](https://github.com/form3tech-oss/chaos-mesh/issues/3716)) ([0bac96e](https://github.com/form3tech-oss/chaos-mesh/commit/0bac96ecb5cea47ebe71b7808455755c536837a2))
* **chaos-dashboard:** ignore the first event on `Reconcile` that triggered by deleting action ([#2698](https://github.com/form3tech-oss/chaos-mesh/issues/2698)) ([857bcc7](https://github.com/form3tech-oss/chaos-mesh/commit/857bcc7645c00cbcd7468ca88ba170dfa3c90739))
* chaos-kernel-builds, python3 and mark bcc version ([#2693](https://github.com/form3tech-oss/chaos-mesh/issues/2693)) ([b23656e](https://github.com/form3tech-oss/chaos-mesh/commit/b23656e886a01e2633831f119c0dcd898ebff208))
* **chaosctl:** best-effort-chaosctl ([#1434](https://github.com/form3tech-oss/chaos-mesh/issues/1434)) ([88e8bad](https://github.com/form3tech-oss/chaos-mesh/commit/88e8bad99bab6f825b62cffa7b9f8e7acc01ae83))
* **chaosctl:** using in-cluster or kubeconfig for port-forwarding ins… ([#1405](https://github.com/form3tech-oss/chaos-mesh/issues/1405)) ([3b4647d](https://github.com/form3tech-oss/chaos-mesh/commit/3b4647d8c7e506edbd43d8cbd0259feeb9a3e0bf))
* chaosfs pseudo random errno UT ([#398](https://github.com/form3tech-oss/chaos-mesh/issues/398)) ([a5519ac](https://github.com/form3tech-oss/chaos-mesh/commit/a5519ac3b568068ac4203109b5d6b2f810d143b6))
* **ci:** resolve part of verify warnings and upgrade revive ([#3776](https://github.com/form3tech-oss/chaos-mesh/issues/3776)) ([191c9ae](https://github.com/form3tech-oss/chaos-mesh/commit/191c9aeae3c3258ab405eff8e3b5e514801b51de))
* **ci:** restrict access to kubeconfig ([#4002](https://github.com/form3tech-oss/chaos-mesh/issues/4002)) ([25a841a](https://github.com/form3tech-oss/chaos-mesh/commit/25a841a23cd2ab3c182fb4f218b8f726c139cf05))
* clear SearchTrigger keyMap eagerly ([#1237](https://github.com/form3tech-oss/chaos-mesh/issues/1237)) ([a10f6c8](https://github.com/form3tech-oss/chaos-mesh/commit/a10f6c8fc6a1c8ccaac099e7b29edf3b7583616d))
* client.Get gets empty GVK(Group/APIVersion/Kind) ([#1368](https://github.com/form3tech-oss/chaos-mesh/issues/1368)) ([39bce28](https://github.com/form3tech-oss/chaos-mesh/commit/39bce28bd191c96e8b8acdbc28b89117b3208661))
* container kill desc ([#3852](https://github.com/form3tech-oss/chaos-mesh/issues/3852)) ([a4ff238](https://github.com/form3tech-oss/chaos-mesh/commit/a4ff23871e7368e0f9f790a4f429ec6a86e364b7))
* correct blog dead link to make lint happy. ([#1077](https://github.com/form3tech-oss/chaos-mesh/issues/1077)) ([1e6f743](https://github.com/form3tech-oss/chaos-mesh/commit/1e6f743e59ae36082ced861e677646a536b7cf5c))
* correct the broken links for get started doc. ([#1073](https://github.com/form3tech-oss/chaos-mesh/issues/1073)) ([a0cab8d](https://github.com/form3tech-oss/chaos-mesh/commit/a0cab8d9882a09cd2720b6c8513567b1f46ade8e))
* correct the usage of pods preview ([#785](https://github.com/form3tech-oss/chaos-mesh/issues/785)) ([645db62](https://github.com/form3tech-oss/chaos-mesh/commit/645db62aae142645bfb8172f569ad3d722ce6911))
* correct update operation in updateStressChaos ([#815](https://github.com/form3tech-oss/chaos-mesh/issues/815)) ([0e3582b](https://github.com/form3tech-oss/chaos-mesh/commit/0e3582bbc27828d31468ce810d58c85586da2c2b))
* **daemon:** allow docker client use low version API ([#2544](https://github.com/form3tech-oss/chaos-mesh/issues/2544)) ([e7bacc1](https://github.com/form3tech-oss/chaos-mesh/commit/e7bacc1f093a32b52080a2f5ab5901c730fd2090))
* **dashboard/core:** double the column size ([#3014](https://github.com/form3tech-oss/chaos-mesh/issues/3014)) ([32dce15](https://github.com/form3tech-oss/chaos-mesh/commit/32dce15f773a686ccc427bfbadb38d33ef745a6f))
* deadline reconciler, keep deadline exceed with true ([#2120](https://github.com/form3tech-oss/chaos-mesh/issues/2120)) ([08facff](https://github.com/form3tech-oss/chaos-mesh/commit/08facff443f3e25459d0c4a41c94af0b57da5ed1))
* disable duration when chaos is instant ([#2269](https://github.com/form3tech-oss/chaos-mesh/issues/2269)) ([de53dce](https://github.com/form3tech-oss/chaos-mesh/commit/de53dce60a1372ff2ec2be8fa364c235257511f4))
* doc versions ([#883](https://github.com/form3tech-oss/chaos-mesh/issues/883)) ([f128602](https://github.com/form3tech-oss/chaos-mesh/commit/f12860223dedb88735d13cc414be142512666471))
* duplicated field created_at with which in gorm.Model ([#2081](https://github.com/form3tech-oss/chaos-mesh/issues/2081)) ([4854fbd](https://github.com/form3tech-oss/chaos-mesh/commit/4854fbdac9595529394e728a0043e2042337b513))
* enable mode when creating PhysicalMachineChaos with addresses ([#3797](https://github.com/form3tech-oss/chaos-mesh/issues/3797)) ([d62eeb4](https://github.com/form3tech-oss/chaos-mesh/commit/d62eeb43331ca4f1065d691c35b817755eea4d7f))
* error keys processing and unset ns selectors ([#1323](https://github.com/form3tech-oss/chaos-mesh/issues/1323)) ([8806c02](https://github.com/form3tech-oss/chaos-mesh/commit/8806c029bf2ce33917beca255b3c503cba08594d))
* experiments ([#582](https://github.com/form3tech-oss/chaos-mesh/issues/582)) ([14493ac](https://github.com/form3tech-oss/chaos-mesh/commit/14493ac9df745ba3b223c541c98e52c76e493879))
* fail to get gitVersion and gitCommit  ([#1377](https://github.com/form3tech-oss/chaos-mesh/issues/1377)) ([1059e15](https://github.com/form3tech-oss/chaos-mesh/commit/1059e153f4ba633d139975787205c8b90268f3a4))
* fetch pods after namespaces changed ([#1178](https://github.com/form3tech-oss/chaos-mesh/issues/1178)) ([4c25dba](https://github.com/form3tech-oss/chaos-mesh/commit/4c25dba525535db4e281c9fb5331d030dedcf477))
* file name ([#3736](https://github.com/form3tech-oss/chaos-mesh/issues/3736)) ([2dad0f6](https://github.com/form3tech-oss/chaos-mesh/commit/2dad0f69640bb4801382f114cf22deb34b3cccdb))
* fix broken CI "Script Test" and "Intergration Test" ([#2067](https://github.com/form3tech-oss/chaos-mesh/issues/2067)) ([e3b091d](https://github.com/form3tech-oss/chaos-mesh/commit/e3b091d0d20c16804473a27052b61a3dbe2af34c))
* fix the bug that event detail can't update ([#1142](https://github.com/form3tech-oss/chaos-mesh/issues/1142)) ([c3bb273](https://github.com/form3tech-oss/chaos-mesh/commit/c3bb2736c5e715c11ea8e0018d96a8297107bae3))
* footer get started link address ([#1019](https://github.com/form3tech-oss/chaos-mesh/issues/1019)) ([71e928a](https://github.com/form3tech-oss/chaos-mesh/commit/71e928ab1673ed098997eaf83ceac3cf58ca43ee))
* gitignore ([#1918](https://github.com/form3tech-oss/chaos-mesh/issues/1918)) ([4aeb6a7](https://github.com/form3tech-oss/chaos-mesh/commit/4aeb6a74d2777cf08f7f428ed10eb7f32a8bcac6))
* handle routes without slash ([#2106](https://github.com/form3tech-oss/chaos-mesh/issues/2106)) ([2bc0fa8](https://github.com/form3tech-oss/chaos-mesh/commit/2bc0fa815446599852116ac0dfb612f659376745))
* helm install error ([#2591](https://github.com/form3tech-oss/chaos-mesh/issues/2591)) ([4caa377](https://github.com/form3tech-oss/chaos-mesh/commit/4caa377251480fe90ad40fdc41a63af225f56106))
* **helm:** ignore entire webhook-configuration if certManager not enabled ([#366](https://github.com/form3tech-oss/chaos-mesh/issues/366)) ([3acc653](https://github.com/form3tech-oss/chaos-mesh/commit/3acc6531ca17285f38c43eaf648ba6143b279224))
* i18n ([#2169](https://github.com/form3tech-oss/chaos-mesh/issues/2169)) ([071e118](https://github.com/form3tech-oss/chaos-mesh/commit/071e118a9a96e4c5f8f1826552f65ccd65d5d61c))
* ignore more unnecessary workflows ([#3738](https://github.com/form3tech-oss/chaos-mesh/issues/3738)) ([ce876cf](https://github.com/form3tech-oss/chaos-mesh/commit/ce876cfa7a94576e77f13530eb8272bb23d0ada6))
* ignore unnecessary workflows when `**.md` or `ui/**` changes ([#3721](https://github.com/form3tech-oss/chaos-mesh/issues/3721)) ([604bef2](https://github.com/form3tech-oss/chaos-mesh/commit/604bef2b02cbe7d0778eb0cdafb7b18780e96406))
* jvmchaos example ([#1944](https://github.com/form3tech-oss/chaos-mesh/issues/1944)) ([5d6e954](https://github.com/form3tech-oss/chaos-mesh/commit/5d6e9541d5a56fb66d3fe7d408f9923d2fb2763e))
* **JVMChaos:** find the correct pid by CommName ([#2904](https://github.com/form3tech-oss/chaos-mesh/issues/2904)) ([68e6ccb](https://github.com/form3tech-oss/chaos-mesh/commit/68e6ccbec79e8d56d2c09b45d4b0c73927d79902))
* key identity problem ([#1040](https://github.com/form3tech-oss/chaos-mesh/issues/1040)) ([c1fd074](https://github.com/form3tech-oss/chaos-mesh/commit/c1fd074b72f397d9a21ea6c16d2bdfe7be522022))
* leader election create cm in target namespace ([#2474](https://github.com/form3tech-oss/chaos-mesh/issues/2474)) ([2a0443f](https://github.com/form3tech-oss/chaos-mesh/commit/2a0443f157a39b7f31cb13df0f910f9acd1ed181))
* let workflow and workflownode bypass the vauth webhook ([#1681](https://github.com/form3tech-oss/chaos-mesh/issues/1681)) ([bc9c86a](https://github.com/form3tech-oss/chaos-mesh/commit/bc9c86a7a43bdbec6554cefce7fe304463e7a27f))
* link of release note ([#982](https://github.com/form3tech-oss/chaos-mesh/issues/982)) ([e4fbe9e](https://github.com/form3tech-oss/chaos-mesh/commit/e4fbe9e389cef4048f119f9ca08d081f57465170))
* links ([#1090](https://github.com/form3tech-oss/chaos-mesh/issues/1090)) ([61457c5](https://github.com/form3tech-oss/chaos-mesh/commit/61457c5fa22aaed8a46980b4f2acb67003c1ba56))
* make restart watcher log level down to info. ([#914](https://github.com/form3tech-oss/chaos-mesh/issues/914)) ([8ce400a](https://github.com/form3tech-oss/chaos-mesh/commit/8ce400a4eb83ece39dbd28d4f5a013850bbf9f84))
* migrate from k8s.gcr.io to registry.k8s.io ([#3974](https://github.com/form3tech-oss/chaos-mesh/issues/3974)) ([de39d6b](https://github.com/form3tech-oss/chaos-mesh/commit/de39d6bc0407dda5ca27213fc57b3db4c3918b1b))
* migration script ([#1983](https://github.com/form3tech-oss/chaos-mesh/issues/1983)) ([b722140](https://github.com/form3tech-oss/chaos-mesh/commit/b722140d915b486f26f4a0171cdf28a57e4238fa))
* **minor:** some ui assets problems ([#1499](https://github.com/form3tech-oss/chaos-mesh/issues/1499)) ([0ed90ce](https://github.com/form3tech-oss/chaos-mesh/commit/0ed90ce78349022cde8677e262d1441bed345638))
* misc ([#3762](https://github.com/form3tech-oss/chaos-mesh/issues/3762)) ([852c402](https://github.com/form3tech-oss/chaos-mesh/commit/852c402243c2d8afe3ae0ee2690e9645365a652c))
* missing append pods in foreach ([#1263](https://github.com/form3tech-oss/chaos-mesh/issues/1263)) ([b05da7d](https://github.com/form3tech-oss/chaos-mesh/commit/b05da7d239e0f56aa8d0079164f498908a3065e8))
* missing created_at field ([#2123](https://github.com/form3tech-oss/chaos-mesh/issues/2123)) ([a457c76](https://github.com/form3tech-oss/chaos-mesh/commit/a457c7626edc312c985c1566157523f71ed602af))
* missing experiment status ([#1144](https://github.com/form3tech-oss/chaos-mesh/issues/1144)) ([e5bc30c](https://github.com/form3tech-oss/chaos-mesh/commit/e5bc30c6f7642fcd7eeda6dafc8752baf05625bb))
* missing PartitionAction ([#1062](https://github.com/form3tech-oss/chaos-mesh/issues/1062)) ([c15f4bd](https://github.com/form3tech-oss/chaos-mesh/commit/c15f4bd2070efb619d62b025cc2f7ba27532d160))
* missing stressors checking ([#1387](https://github.com/form3tech-oss/chaos-mesh/issues/1387)) ([f6e0637](https://github.com/form3tech-oss/chaos-mesh/commit/f6e06375f6517f6020a5cfcba2476686a14b0daa))
* Omit optional fields explicitly ([#3531](https://github.com/form3tech-oss/chaos-mesh/issues/3531)) ([626d379](https://github.com/form3tech-oss/chaos-mesh/commit/626d37957e1c3c016890134770cd6c6057604d12))
* only schedule mode in pod(container)-kill ([#597](https://github.com/form3tech-oss/chaos-mesh/issues/597)) ([c7003eb](https://github.com/form3tech-oss/chaos-mesh/commit/c7003eb7a998da0b0706185910411a243e5b2326))
* optimize darwin static check ([#3570](https://github.com/form3tech-oss/chaos-mesh/issues/3570)) ([62f9e43](https://github.com/form3tech-oss/chaos-mesh/commit/62f9e43c6bb1a32ae19b0dba07da626fc17dce85))
* pass correct value to scope mode fixed ([#814](https://github.com/form3tech-oss/chaos-mesh/issues/814)) ([3223b44](https://github.com/form3tech-oss/chaos-mesh/commit/3223b44e0ec432067e3fe2f5f5e1e8e84333e45f))
* persist event with UTC Time ([#2303](https://github.com/form3tech-oss/chaos-mesh/issues/2303)) ([1b8447a](https://github.com/form3tech-oss/chaos-mesh/commit/1b8447ad561c0ff4cd86fcbd688bee7f3e61058a))
* phase_selectors to phaseSelectors ([#2330](https://github.com/form3tech-oss/chaos-mesh/issues/2330)) ([3d918f1](https://github.com/form3tech-oss/chaos-mesh/commit/3d918f1623968f61de87c257c2a320ed86c8129b))
* pod deletion caused twophase phase=Failed chaos. ([#946](https://github.com/form3tech-oss/chaos-mesh/issues/946)) ([44e32fe](https://github.com/form3tech-oss/chaos-mesh/commit/44e32fe77a3d55a882e689171e0d5ebeaf29e554))
* pr template ([#3571](https://github.com/form3tech-oss/chaos-mesh/issues/3571)) ([88eede8](https://github.com/form3tech-oss/chaos-mesh/commit/88eede8c3b7c292f57225fdc4213c422cf37bc78))
* **readme:** update crd types ([#1201](https://github.com/form3tech-oss/chaos-mesh/issues/1201)) ([7d80ca5](https://github.com/form3tech-oss/chaos-mesh/commit/7d80ca5e8556c86a4dbd00578aba6018a27dc35c))
* refactor service level in pkg/dashboard/apiserver/experiment ([#3445](https://github.com/form3tech-oss/chaos-mesh/issues/3445)) ([d8e900b](https://github.com/form3tech-oss/chaos-mesh/commit/d8e900bff3b6cec53c462b95899a4e176289aa09))
* relative paths' problem when first rendering ([#675](https://github.com/form3tech-oss/chaos-mesh/issues/675)) ([c463294](https://github.com/form3tech-oss/chaos-mesh/commit/c463294d7b6cdf6a353f757dd5f323f6faa5b9f2))
* remove empty priorityClassName ([#2863](https://github.com/form3tech-oss/chaos-mesh/issues/2863)) ([b3f298e](https://github.com/form3tech-oss/chaos-mesh/commit/b3f298e9a1a0cafc2209a5e06b30f23ecc27c0fe))
* remove hardcoded ports and expose them on helm level ([#2571](https://github.com/form3tech-oss/chaos-mesh/issues/2571)) ([d66d1c6](https://github.com/form3tech-oss/chaos-mesh/commit/d66d1c64589c4d77d1d515cd40c7ef4f6cd0bccc))
* remove limit action from BlockChaos ([#3655](https://github.com/form3tech-oss/chaos-mesh/issues/3655)) ([9fc0fad](https://github.com/form3tech-oss/chaos-mesh/commit/9fc0fad31cde59cafaca4bc357067257bed9221c))
* remove rbac deprecation ([#994](https://github.com/form3tech-oss/chaos-mesh/issues/994)) ([981f4e7](https://github.com/form3tech-oss/chaos-mesh/commit/981f4e7f26aa41d9e853da0127959bea0850eec0))
* remove the explicit use of pingcap/log ([#3674](https://github.com/form3tech-oss/chaos-mesh/issues/3674)) ([ed62c84](https://github.com/form3tech-oss/chaos-mesh/commit/ed62c84a51e5f30733dd1318eda2513eca102fe9))
* remove unused stressor in StressChaos ([#753](https://github.com/form3tech-oss/chaos-mesh/issues/753)) ([64522bf](https://github.com/form3tech-oss/chaos-mesh/commit/64522bfbdd9e49a4b3e8e515f489d87003966d9c))
* repair field names to camelCase ([#2408](https://github.com/form3tech-oss/chaos-mesh/issues/2408)) ([e96f33a](https://github.com/form3tech-oss/chaos-mesh/commit/e96f33a07a600181b8d6ee875890aef74b0e35e4))
* replace actions/checkout from master to v3 ([#3737](https://github.com/form3tech-oss/chaos-mesh/issues/3737)) ([ece0909](https://github.com/form3tech-oss/chaos-mesh/commit/ece0909b646754fee37a7df158a31e4a780bf3f2))
* replace core.SelectorInfo with v1alpha1.PodSelectorSpec ([#2364](https://github.com/form3tech-oss/chaos-mesh/issues/2364)) ([91c4914](https://github.com/form3tech-oss/chaos-mesh/commit/91c4914dac9a58546315bb365cd0e7159ec13df2))
* replace env variables with contexts ([#3773](https://github.com/form3tech-oss/chaos-mesh/issues/3773)) ([9d17421](https://github.com/form3tech-oss/chaos-mesh/commit/9d174215eed80794b0fd13cf9c655b23ccef55a0))
* replaced localhost:5000 with ghcr.io as registry name ([#2919](https://github.com/form3tech-oss/chaos-mesh/issues/2919)) ([07508be](https://github.com/form3tech-oss/chaos-mesh/commit/07508bed42066ab1dc22993fbbb0a62978c66f9a))
* **route:** batch delete schedule archives ([#1957](https://github.com/form3tech-oss/chaos-mesh/issues/1957)) ([436cf80](https://github.com/form3tech-oss/chaos-mesh/commit/436cf80fa5bafbd21c300ffc0f87cf2edff6edfb))
* sanitize generated object ([#2511](https://github.com/form3tech-oss/chaos-mesh/issues/2511)) ([3cb9d54](https://github.com/form3tech-oss/chaos-mesh/commit/3cb9d54f8bbbe0f006b6104887ec62299958df26))
* save container_name ([#603](https://github.com/form3tech-oss/chaos-mesh/issues/603)) ([441f1e2](https://github.com/form3tech-oss/chaos-mesh/commit/441f1e21f036646c9247802c73eca022a76e4c51))
* scope and scheduler bugs ([#1166](https://github.com/form3tech-oss/chaos-mesh/issues/1166)) ([0a5e5d4](https://github.com/form3tech-oss/chaos-mesh/commit/0a5e5d4197058d204e3f44f466ab18cceb505c5f))
* scope pods' isolation problem ([#1386](https://github.com/form3tech-oss/chaos-mesh/issues/1386)) ([36bb4f0](https://github.com/form3tech-oss/chaos-mesh/commit/36bb4f013eb8ffd8388023b35f3bd8a9fe69f356))
* scope selectors ([#881](https://github.com/form3tech-oss/chaos-mesh/issues/881)) ([293af7a](https://github.com/form3tech-oss/chaos-mesh/commit/293af7a6fb8e697968bbc0f3ceb8ee75de7195cf))
* set `replicas: 1` automatically when HA is not enabled ([#4079](https://github.com/form3tech-oss/chaos-mesh/issues/4079)) ([f8d938d](https://github.com/form3tech-oss/chaos-mesh/commit/f8d938d6e63c245ebd27a82b2000775800e8fed7))
* set certmanager ca duration ([#3128](https://github.com/form3tech-oss/chaos-mesh/issues/3128)) ([30c36f6](https://github.com/form3tech-oss/chaos-mesh/commit/30c36f625da1df154ff8c6465fd543b743f535ed))
* some bug fixes related to the UI ([#1430](https://github.com/form3tech-oss/chaos-mesh/issues/1430)) ([428ad16](https://github.com/form3tech-oss/chaos-mesh/commit/428ad16c0e75e1f72152bb4782b4dd0422dc03ac))
* some bugs ([#828](https://github.com/form3tech-oss/chaos-mesh/issues/828)) ([84208f3](https://github.com/form3tech-oss/chaos-mesh/commit/84208f3043200802b916ab48b527736a7cb53720))
* some errors in install.sh ([#2858](https://github.com/form3tech-oss/chaos-mesh/issues/2858)) ([712f58c](https://github.com/form3tech-oss/chaos-mesh/commit/712f58cadb6f648fb7d27cf3cabde24dbe1b83b1))
* some minor problems ([#786](https://github.com/form3tech-oss/chaos-mesh/issues/786)) ([8251326](https://github.com/form3tech-oss/chaos-mesh/commit/8251326fddbf1cc60580da6b0b4d392f4c45f734))
* some UI bugs ([#1369](https://github.com/form3tech-oss/chaos-mesh/issues/1369)) ([cb2f3de](https://github.com/form3tech-oss/chaos-mesh/commit/cb2f3de334557699d31199f9234f8c15e2e1876f))
* some ui issues ([#2680](https://github.com/form3tech-oss/chaos-mesh/issues/2680)) ([4bcaeac](https://github.com/form3tech-oss/chaos-mesh/commit/4bcaeac7b5eb44d00ffb56885464f7a9c5f126d9))
* startup workflow bootstrap with fx ([#1929](https://github.com/form3tech-oss/chaos-mesh/issues/1929)) ([4591eac](https://github.com/form3tech-oss/chaos-mesh/commit/4591eac4006d79a7bd9238b6dd448268a1f55fe3))
* supplement mapstructure renaming ([#1212](https://github.com/form3tech-oss/chaos-mesh/issues/1212)) ([eebc349](https://github.com/form3tech-oss/chaos-mesh/commit/eebc349570840999262fad436e8a1d400b2061e5))
* supplement read-only token ([#3971](https://github.com/form3tech-oss/chaos-mesh/issues/3971)) ([d3c8a29](https://github.com/form3tech-oss/chaos-mesh/commit/d3c8a29344540c41472253395b4548dd4a314599))
* support aws/gcp chaos in dashboard archive ([#2337](https://github.com/form3tech-oss/chaos-mesh/issues/2337)) ([7f650e6](https://github.com/form3tech-oss/chaos-mesh/commit/7f650e692cb749b0d474ffd899c4034fed553e5c))
* swagger params missing ([#647](https://github.com/form3tech-oss/chaos-mesh/issues/647)) ([2a1d6b6](https://github.com/form3tech-oss/chaos-mesh/commit/2a1d6b61d3d24cdc3e781ca2ca839a18181016ec))
* **swagger:** update `is mandatory` to true ([#3743](https://github.com/form3tech-oss/chaos-mesh/issues/3743)) ([8b13f5a](https://github.com/form3tech-oss/chaos-mesh/commit/8b13f5ad252f4b702f691480266404d14734549e))
* sync PhysicalMachineChaos to API client and forms ([#3660](https://github.com/form3tech-oss/chaos-mesh/issues/3660)) ([a6620f7](https://github.com/form3tech-oss/chaos-mesh/commit/a6620f7d6c5297d90f958e2e4054743daa7a6a92))
* task node could also be trigger by child workflow node ([#2044](https://github.com/form3tech-oss/chaos-mesh/issues/2044)) ([91ed9b6](https://github.com/form3tech-oss/chaos-mesh/commit/91ed9b6dd317bf0285df5f2c32d2460b7493e112))
* the bad camelCase conversion method ([#1937](https://github.com/form3tech-oss/chaos-mesh/issues/1937)) ([128651f](https://github.com/form3tech-oss/chaos-mesh/commit/128651fbfed8e68b2d9561da25bf44120f7129eb))
* the broken webhook for workflow ([#2225](https://github.com/form3tech-oss/chaos-mesh/issues/2225)) ([984b3b3](https://github.com/form3tech-oss/chaos-mesh/commit/984b3b3140e9f9b11959c8ae922a66d5f6ab26cf))
* the download url of kubebuilder ([#2249](https://github.com/form3tech-oss/chaos-mesh/issues/2249)) ([6bc1166](https://github.com/form3tech-oss/chaos-mesh/commit/6bc11669fc72171cb57e479b25476b1842ea8f71))
* **time chaos:** timechaos not injected to the child process ([#3725](https://github.com/form3tech-oss/chaos-mesh/issues/3725)) ([f45d612](https://github.com/form3tech-oss/chaos-mesh/commit/f45d612f6961d61456345a909eac9fa38898bc65))
* tiny errors in frontend ([#960](https://github.com/form3tech-oss/chaos-mesh/issues/960)) ([18ac4b4](https://github.com/form3tech-oss/chaos-mesh/commit/18ac4b498839afbad069c84f667c13accbb60997))
* tiny typo ([#1131](https://github.com/form3tech-oss/chaos-mesh/issues/1131)) ([41a963e](https://github.com/form3tech-oss/chaos-mesh/commit/41a963e9e22bc9a5d5b2c6e682a81b85ab791cfb))
* token name duplicate ([#1644](https://github.com/form3tech-oss/chaos-mesh/issues/1644)) ([c542178](https://github.com/form3tech-oss/chaos-mesh/commit/c54217863369e9da1015de443d748035775c7243))
* trying to fix real_gettimeofday on arm64 ([#2849](https://github.com/form3tech-oss/chaos-mesh/issues/2849)) ([4d3bfed](https://github.com/form3tech-oss/chaos-mesh/commit/4d3bfeda0f2c68c0416cfbf06c132191f751af6a))
* typo and format pr template ([#617](https://github.com/form3tech-oss/chaos-mesh/issues/617)) ([4fb903d](https://github.com/form3tech-oss/chaos-mesh/commit/4fb903df009b781ecfed1460c2a149ece97ffd5c))
* **ui:** `space` to `enter` in helper texts ([#3335](https://github.com/form3tech-oss/chaos-mesh/issues/3335)) ([89e4cb1](https://github.com/form3tech-oss/chaos-mesh/commit/89e4cb193b51cd4d1e3cc0cfb3e85fb4e53292e9))
* **ui:** IOChaos `containerNames` field ([#3533](https://github.com/form3tech-oss/chaos-mesh/issues/3533)) ([5c806cd](https://github.com/form3tech-oss/chaos-mesh/commit/5c806cd18b6e197d3c4a1659f342423d299c0160))
* **ui:** isolate scopes when creating NetworkChaos ([#3223](https://github.com/form3tech-oss/chaos-mesh/issues/3223)) ([6111432](https://github.com/form3tech-oss/chaos-mesh/commit/61114328c4dffdb9b204658987c1d9dbf7930108))
* **ui:** pod phases should be first letter capitalized ([#2915](https://github.com/form3tech-oss/chaos-mesh/issues/2915)) ([2f68637](https://github.com/form3tech-oss/chaos-mesh/commit/2f68637590468dfd9cc209df22cbf8c7b49234c8))
* **ui:** unable to load from objects ([#2585](https://github.com/form3tech-oss/chaos-mesh/issues/2585)) ([5cdb20f](https://github.com/form3tech-oss/chaos-mesh/commit/5cdb20fd0e0be39f0739eb7c68a184c8370e2df0))
* unhandled empty value in LabelField ([#1468](https://github.com/form3tech-oss/chaos-mesh/issues/1468)) ([3d1a9ee](https://github.com/form3tech-oss/chaos-mesh/commit/3d1a9ee2ac4708e23b6e4117ed2f2074df7f5ce2))
* upgrade pnpm to v8 ([#4051](https://github.com/form3tech-oss/chaos-mesh/issues/4051)) ([afab5b4](https://github.com/form3tech-oss/chaos-mesh/commit/afab5b4a781304ea7e7e7ba1e01262cc3a9f12c2))
* upload predefined objects on UI ([#2068](https://github.com/form3tech-oss/chaos-mesh/issues/2068)) ([e058a08](https://github.com/form3tech-oss/chaos-mesh/commit/e058a08c076e65e76a1455d942d1dbd5ccb24f6e))
* use `*time.Time` to represent finish time ([#4056](https://github.com/form3tech-oss/chaos-mesh/issues/4056)) ([2dc55a2](https://github.com/form3tech-oss/chaos-mesh/commit/2dc55a234727d08613ac1f5b9766c1523967c793))
* use emoji directly in issue config ([#3667](https://github.com/form3tech-oss/chaos-mesh/issues/3667)) ([010070f](https://github.com/form3tech-oss/chaos-mesh/commit/010070f70010d6cb4f34f1bd29ca8afc7a678fba))
* use patch instead of update within podfailure ([#2020](https://github.com/form3tech-oss/chaos-mesh/issues/2020)) ([bd09176](https://github.com/form3tech-oss/chaos-mesh/commit/bd091764f02df5a5cae12e06ee2b9664c3c502f4))
* validation problem in dashboard ([#968](https://github.com/form3tech-oss/chaos-mesh/issues/968)) ([83a2175](https://github.com/form3tech-oss/chaos-mesh/commit/83a217530481b32d9f8528b59f0bd548bc438009))
* **workflow:** no more event after accomplished ([#2911](https://github.com/form3tech-oss/chaos-mesh/issues/2911)) ([8523568](https://github.com/form3tech-oss/chaos-mesh/commit/852356876103b4f0269156342137da385c13a915))

## [2.6.0] - 2023-05-30

### Added

- Install offline Helm Chart for a multi-cluster [#3897](https://github.com/chaos-mesh/chaos-mesh/pull/3897)

### Changed

- Change CoreDNS listen port from 53 to 5353 [#4022](https://github.com/chaos-mesh/chaos-mesh/pull/4022)
- Bump go to v1.19.3 [#3770](https://github.com/chaos-mesh/chaos-mesh/pull/3770)
- Change ubuntu version from latest to 20.04 [#3817](https://github.com/chaos-mesh/chaos-mesh/pull/3817)
- Switch views between k8s and hosts nodes [#3830](https://github.com/chaos-mesh/chaos-mesh/pull/3830)
- New CI for finding merge conflicts [#3850](https://github.com/chaos-mesh/chaos-mesh/pull/3850)
- Upgrade byteman-helper to v4.0.20 [#3863](https://github.com/chaos-mesh/chaos-mesh/pull/3863)
- Helm: change default webhook port to 10250 [#3877](https://github.com/chaos-mesh/chaos-mesh/pull/3877)
- Upgrade base image for chaos-mesh to alpine:3.17 [#3893](https://github.com/chaos-mesh/chaos-mesh/pull/3893)
- Slow down releasing the latest version [#3900](https://github.com/chaos-mesh/chaos-mesh/pull/3900)
- Update k8s.io dependencies to v0.26.1 [#3902](https://github.com/chaos-mesh/chaos-mesh/pull/3902)
- Update sigs.k8s.io/controller-runtime to v0.14.1 and sigs.k8s.io/controller-tools to v0.11.1 [#3902](https://github.com/chaos-mesh/chaos-mesh/pull/3902)
- Change the package manager from `yarn` to `pnpm`. [#3965](https://github.com/chaos-mesh/chaos-mesh/pull/3965)
- Upgrade DNS CoreDNS image url to ghcr.io [#3488](https://github.com/chaos-mesh/chaos-mesh/pull/3488)
- Upgrade OS image for chaos-daemon container image [#3905](https://github.com/chaos-mesh/chaos-mesh/pull/3905)
- Replace openapi-generator with Orval and React Query [#3748](https://github.com/chaos-mesh/chaos-mesh/pull/3748)
- Cleanup makefile and provide `make help` [#3888](https://github.com/chaos-mesh/chaos-mesh/pull/3888)
- Remove `IN_DOCKER` environment variable in `Makefile` [#3992](https://github.com/chaos-mesh/chaos-mesh/pull/3992)
- Refine TTL config of Chaos dashboard [#4008](https://github.com/chaos-mesh/chaos-mesh/pull/4008)
- `pause` would return non zero exit code when the subcommand failed [#4018](https://github.com/chaos-mesh/chaos-mesh/pull/4018)
- use helm values to set chaos-daemon capabilities [#4030](https://github.com/chaos-mesh/chaos-mesh/pull/4030)
- Build binaries locally with `local/` prefix targets in `Makefile` [#4004](https://github.com/chaos-mesh/chaos-mesh/pull/4004)
- Use kubectl cluster-info dump to enhance e2e profiling [#3759](https://github.com/chaos-mesh/chaos-mesh/pull/3759)
- Upgrade fx event logger [#4036](https://github.com/chaos-mesh/chaos-mesh/pull/4036)
- Refine logging in `pkg/selector/physicalmachine` [#4037](https://github.com/chaos-mesh/chaos-mesh/pull/4037)
- Setup OWNERS and OWNERS_ALIASES [#4039](https://github.com/chaos-mesh/chaos-mesh/pull/4039)

### Deprecated

- Nothing

### Removed

- Remove no needed file crd-v1beta1.yaml [#3807](https://github.com/chaos-mesh/chaos-mesh/pull/3807)
- Remove useless kubebuilder comment in webhook [#3816](https://github.com/chaos-mesh/chaos-mesh/pull/3816)
- Remove the unused inject-v1-pod webhook. [#3885](https://github.com/chaos-mesh/chaos-mesh/pull/3885)

### Fixed

- Fix version comparison in install.sh [#3901](https://github.com/chaos-mesh/chaos-mesh/pull/3901)
- Fix stuck dashboard updates when using ReadWriteOnce PVCs [#3876](https://github.com/chaos-mesh/chaos-mesh/issues/3876)
- Fix MySQL NO_ZERO_IN_DATE by using `*time.Time` to represent finish time [#4056](https://github.com/chaos-mesh/chaos-mesh/pull/4056)

### Security

- Bump go to v1.19.7 to fix CVE-2022-41723 [#3978](https://github.com/chaos-mesh/chaos-mesh/pull/3978) [#3981](https://github.com/chaos-mesh/chaos-mesh/pull/3981)
- Bump github.com/opencontainers/runc from 1.1.4 to 1.1.5 [#3987](https://github.com/chaos-mesh/chaos-mesh/pull/3987)

## [2.5.0] - 2022-11-22

### Added

- Add `controller.securityContext` and `dashboard.securityContext` to Helm chart [#3603](https://github.com/chaos-mesh/chaos-mesh/pull/3603)
- Add `RemoteCluster` resource type [#3342](https://github.com/chaos-mesh/chaos-mesh/pull/3342)
- Add `clusterregistry` package to help developers to develop multi-cluster reconciler [#3342](https://github.com/chaos-mesh/chaos-mesh/pull/3342)
- Add features about integration with helm to install Chaos Mesh in remote cluster [#3384](https://github.com/chaos-mesh/chaos-mesh/pull/3384)
- Add new CI "Manually Sign Container Images" to sign existing container images [#3708](https://github.com/chaos-mesh/chaos-mesh/pull/3708)
- Install and uninstall chaos mesh in remote cluster through `RemoteCluster` resource [#3414](https://github.com/chaos-mesh/chaos-mesh/pull/3414)
- MultiCluster: support inject / recover on remote cluster [#3453](https://github.com/chaos-mesh/chaos-mesh/pull/3453)
- Add TLS support for HTTPChaos [#3549](https://github.com/chaos-mesh/chaos-mesh/pull/3549)

### Changed

- Use the next generation `New Workflow` UI by default [#3718](https://github.com/chaos-mesh/chaos-mesh/pull/3718)
- StressChaos: Support cgroup v2 for docker and crio [#3698](https://github.com/chaos-mesh/chaos-mesh/pull/3698)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Remove the explicit use of pingcap/log [#3674](https://github.com/chaos-mesh/chaos-mesh/pull/3674)
- Fix typo in controller error message [#3704](https://github.com/chaos-mesh/chaos-mesh/pull/3704)
- Fix panic when logging, log kvs as pair [#3716](https://github.com/chaos-mesh/chaos-mesh/pull/3716)
- Fix timechaos not injected into the child process [#3725](https://github.com/chaos-mesh/chaos-mesh/pull/3725)
- Ignore `ScheduleSkipRemoveHistory` events to fix the memory of controller-manager keep increasing [#3761](https://github.com/chaos-mesh/chaos-mesh/issues/3761)
- Update `is mandatory` to true in a swagger comment [#3743](https://github.com/chaos-mesh/chaos-mesh/pull/3743)
- Enable mode when creating PhysicalMachineChaos with addresses in UI [#3797](https://github.com/chaos-mesh/chaos-mesh/pull/3797)

### Security

- Sign images and generate sbom when uploading images in CI [#3766](https://github.com/chaos-mesh/chaos-mesh/pull/3766)

## [2.4.2] - 2022-11-07

### Added

- Nothing

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix timechaos not injected into the child process [#3730](https://github.com/chaos-mesh/chaos-mesh/pull/3730)
- Update `is mandatory` to true in a swagger comment [#3743](https://github.com/chaos-mesh/chaos-mesh/pull/3743)

### Security

- Nothing

## [2.4.1] - 2022-09-27

### Added

- Nothing

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix: mark 2.4.1 not as pre-release version[#3679](https://github.com/chaos-mesh/chaos-mesh/pull/3679)

### Security

- Nothing

## [2.4.0] - 2022-09-23

### Added

- Add support for `PhysicalMachine` in UI [#3624](https://github.com/chaos-mesh/chaos-mesh/pull/3624)

### Changed

- Replace io/ioutil package with os package. [#3539](https://github.com/chaos-mesh/chaos-mesh/pull/3539)
- Refine logging in pkg/dashboard/apiserver/event, moved package level variable log into struct Service as a field. [#3528](https://github.com/chaos-mesh/chaos-mesh/pull/3528)
- Refine logging in pkg/dashboard/apiserver/auth/gcp, moved package level variable log into struct Service as a field [#3527](https://github.com/chaos-mesh/chaos-mesh/pull/3527)
- Change e2e config settings to append "pause image" args. [#3567](https://github.com/chaos-mesh/chaos-mesh/pull/3567)
- Update display of disabled scope in UI [#3621](https://github.com/chaos-mesh/chaos-mesh/pull/3621)
- Make the Scope render conditionally [#3622](https://github.com/chaos-mesh/chaos-mesh/pull/3622)
- Refine logging, remove the usage of klog in chaosctl [#3628](https://github.com/chaos-mesh/chaos-mesh/pull/3628)
- Dashboard: Fix rbac.yaml for token generation verbs/resources [#3370](https://github.com/chaos-mesh/chaos-mesh/pull/3370)
- Refine logging, remove the usage of klog in event recorder [#3629](https://github.com/chaos-mesh/chaos-mesh/pull/3629)
- Use int64 to restore latency for BlockChaos [#3638](https://github.com/chaos-mesh/chaos-mesh/pull/3638)
- Remove CRD v1beta1 support [#3630](https://github.com/chaos-mesh/chaos-mesh/pull/3630)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix NetworkChaos fail with @ifXX in the device [#3605](https://github.com/chaos-mesh/chaos-mesh/pull/3605)
- Fix BlockChaos can't show Chinese name. [#3536](https://github.com/chaos-mesh/chaos-mesh/pull/3536)
- Add `omitempty` JSON tag to optional fields of the CRD objects. [#3531](https://github.com/chaos-mesh/chaos-mesh/pull/3531)
- Fix "sidecar config" e2e test cases run failed in some scenario.[#3564](https://github.com/chaos-mesh/chaos-mesh/pull/3564)
- Fix Integration test with bumping kubectl version. [#3589](https://github.com/chaos-mesh/chaos-mesh/pull/3589)

### Security

- Nothing

## [2.3.3] - 2022-11-07

### Added

- Nothing

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Remove `limit` action of `BlockChaos` in the dashboard [#3655](https://github.com/chaos-mesh/chaos-mesh/pull/3655)
- Sync PhysicalMachineChaos to API client and forms [#3660](https://github.com/chaos-mesh/chaos-mesh/pull/3660)
- Fix timechaos not injected into the child process [#3729](https://github.com/chaos-mesh/chaos-mesh/pull/3729)
- Update `is mandatory` to true in a swagger comment [#3743](https://github.com/chaos-mesh/chaos-mesh/pull/3743)

### Security

- Nothing

## [2.3.2] - 2022-09-20

### Added

- Nothing

### Changed

- Use int64 to restore latency for BlockChaos [#3638](https://github.com/chaos-mesh/chaos-mesh/pull/3638)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix: update helm chart annotation artifacthub.io/prerelease to false [#3626](https://github.com/chaos-mesh/chaos-mesh/pull/3626)

### Security

- Nothing

## [2.3.1] - 2022-09-02

### Added

- Nothing

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Protect chaos available namespace and filter namespaces if needed [#3473](https://github.com/chaos-mesh/chaos-mesh/pull/3473)
- Respect flag `enableProfiling` and do not register profiler endpoints when it's false [#3474](https://github.com/chaos-mesh/chaos-mesh/pull/3474)
- Fix the blank screen after creating chaos experiment with "By YAML" [#3489](https://github.com/chaos-mesh/chaos-mesh/pull/3489)
- Update hint text about the manual token generating process for Kubernetes 1.24+ [#3505](https://github.com/chaos-mesh/chaos-mesh/pull/3505)
- Fix IOChaos `containerNames` field in UI [#3533](https://github.com/chaos-mesh/chaos-mesh/pull/3533)
- Fix BlockChaos can't show Chinese name. [#3536](https://github.com/chaos-mesh/chaos-mesh/pull/3536)
- Fix recover bug when setting force recover to true [#3578](https://github.com/chaos-mesh/chaos-mesh/pull/3578)
- Fix generate token failed on chaos dashboard [#3595](https://github.com/chaos-mesh/chaos-mesh/pull/3595)

### Security

- Nothing

## [2.3.0] - 2022-07-29

### Added

- Add more status for record [#3170](https://github.com/chaos-mesh/chaos-mesh/pull/3170)
- Add `chaosDaemon.updateStrategy` to Helm chart to allow configuring `DaemonSetUpdateStrategy` for chaos-daemon [#3108](https://github.com/chaos-mesh/chaos-mesh/pull/3108)
- Add AArch64 support for TimeChaos [#3088](https://github.com/chaos-mesh/chaos-mesh/pull/3088)
- Add integration test and link test on arm [#3177](https://github.com/chaos-mesh/chaos-mesh/pull/3177)
- Add `spec.privateKey.rotationPolicy` to Certificates, to comply with requirements in cert-manager 1.8 [#3325](https://github.com/chaos-mesh/chaos-mesh/pull/3325)
- Add `clusterregistry` package to help developers to develop multi-cluster reconciler [#3342](https://github.com/chaos-mesh/chaos-mesh/pull/3342)
- Support `Suspend` in next generation `New Workflow`'s UI [#3254](https://github.com/chaos-mesh/chaos-mesh/pull/3254)
- Add helm annotations for Artifact Hub [#3355](https://github.com/chaos-mesh/chaos-mesh/pull/3355)
- Add implementation of blockchaos in chaos-daemon [#2907](https://github.com/chaos-mesh/chaos-mesh/pull/2907)
- Bump chaos-tproxy to v0.5.1 [#3412](https://github.com/chaos-mesh/chaos-mesh/pull/3412)
- Allow importing external workflows and copying flow nodes in next generation `New Workflow` [#3368](https://github.com/chaos-mesh/chaos-mesh/pull/3368)
- Add `QPS` and `Burst` for Chaos Dashboard Configuration [#3476](https://github.com/chaos-mesh/chaos-mesh/pull/3476)
- Add guide and example for monitoring Chaos Mesh [#3030](https://github.com/chaos-mesh/chaos-mesh/pull/3030)
- Support `KernelChaos` in `AutoForm` [#3449](https://github.com/chaos-mesh/chaos-mesh/pull/3449)
- Sync latest Chaosd and PhysicalMachineChaos [#3477](https://github.com/chaos-mesh/chaos-mesh/pull/3477)
- Add accept-tcp-flag to network delay in PysicalMachineChaos [#3588](https://github.com/chaos-mesh/chaos-mesh/pull/3588)

### Changed

- Helm charts: update validate-auth to chaos-mesh-validation-auth [#3193](https://github.com/chaos-mesh/chaos-mesh/pull/3193)
- Helm charts: support latest api version of dashboard ingress [#3066](https://github.com/chaos-mesh/chaos-mesh/pull/3066)
- Update shell script to support shellchecks [#3230](https://github.com/chaos-mesh/chaos-mesh/pull/3230)
- CI: build dev-env and build-env for e2e tests if required [#3264](https://github.com/chaos-mesh/chaos-mesh/pull/3264)
- CI: version unrelated manifests [#3293](https://github.com/chaos-mesh/chaos-mesh/pull/3293)
- Bump chaos-tproxy to v0.4.6 [#3272](https://github.com/chaos-mesh/chaos-mesh/pull/3272)
- Helm charts: using 0.0.0 as version and appVersion [#3311](https://github.com/chaos-mesh/chaos-mesh/pull/3311)
- Add a comment to the flag size of memory stress in the dashboard [#3359](https://github.com/chaos-mesh/chaos-mesh/pull/3359)
- Refine logging in pkg/dashboard/store, removed global the log [#3143](https://github.com/chaos-mesh/chaos-mesh/pull/3143)
- Renamed namespace from chaos-testing to chaos-mesh [#3353](https://github.com/chaos-mesh/chaos-mesh/pull/3353)
- Use ContainerSelector in kernel chaos [#3395](https://github.com/chaos-mesh/chaos-mesh/pull/3395)
- Make possible to have more than one dns chaos server [#3381](https://github.com/chaos-mesh/chaos-mesh/pull/3381)
- Helm charts: Relax allowedHostPaths in chaos-daemon PSP [#3350](https://github.com/chaos-mesh/chaos-mesh/pull/3350)
- Run build image ci on self-hosted machine [#3429](https://github.com/chaos-mesh/chaos-mesh/pull/3429)
- Simplified logic and add test case about finalizers. [#3422](https://github.com/chaos-mesh/chaos-mesh/pull/3422)
- Update API requests with OpenAPI generated client [#2926](https://github.com/chaos-mesh/chaos-mesh/pull/2926)
- Implement some missing methods in ctrl server [#3462](https://github.com/chaos-mesh/chaos-mesh/pull/3462)
- Use `net.Interfaces()` to implement `getAllInterfaces()` [#3484](https://github.com/chaos-mesh/chaos-mesh/pull/3484)

### Deprecated

- Nothing

### Removed

- Removed extra import of common pkg in chaosctl/cmd/logs.go
- Removed unused local function from statuscheck/manager.go [#3228](https://github.com/chaos-mesh/chaos-mesh/pull/3228)
- Removed ui build and test for arm64 [#3305](https://github.com/chaos-mesh/chaos-mesh/pull/3305)
- Remove sed need (SC2001) [#3248](https://github.com/chaos-mesh/chaos-mesh/pull/3248)
- Removed not used clientset in cmd/chaos-controller-manager/main.go [#3334](https://github.com/chaos-mesh/chaos-mesh/pull/3334)
- Removed not used globalCacheReader in cmd/chaos-controller-manager/provider/controller.go [#3343](https://github.com/chaos-mesh/chaos-mesh/pull/3343)
- Removed unsupported action comments of blockchaos [#3435](https://github.com/chaos-mesh/chaos-mesh/pull/3435)

### Fixed

- Update description of memory stressors [#3225](https://github.com/chaos-mesh/chaos-mesh/pull/3225)
- Isolate `target` field and `Scope` when creating `NetworkChaos` in UI [#3223](https://github.com/chaos-mesh/chaos-mesh/issues/3221)
- Add arm64 architecture to ci_skip to pass required test [#3305](https://github.com/chaos-mesh/chaos-mesh/pull/3305)
- Adapt install.sh for kubectl/kubernetes cluster greater than 1.24 [#3177](https://github.com/chaos-mesh/chaos-mesh/pull/3177)
- SC2166: Use || or && rather than -o or -a [#3235](https://github.com/chaos-mesh/chaos-mesh/pull/3235)
- SC2206: Use quote to prevent word splitting/globbing [#3234](https://github.com/chaos-mesh/chaos-mesh/pull/3234)
- Fix make check does not respect the env-images.yaml [#3210] (<https://github.com/chaos-mesh/chaos-mesh/pull/3210>)
- SC2004: Remove unnecessary $ on arithmetic variables [#3247](https://github.com/chaos-mesh/chaos-mesh/pull/3247)
- PhysicalMachineChaos: update stress options type [#3347](https://github.com/chaos-mesh/chaos-mesh/pull/3347)
- PhysicalMachineChaos: remove validate for IP and host for delay, loss, duplicate, corruption [#3483](https://github.com/chaos-mesh/chaos-mesh/pull/3483)
- StressChaos: run `pause` before `choom` [#3405](https://github.com/chaos-mesh/chaos-mesh/pull/3405)
- JVMChaos: update the error message that can be ignored [#3415](https://github.com/chaos-mesh/chaos-mesh/pull/3415)
- Fix Workflow Validating Webhook Panic [#3413](https://github.com/chaos-mesh/chaos-mesh/pull/3413)
- Overwrite $IMAGE_BUILD_ENV_TAG with $IMAGE_TAG-$ARCH in `upload_env_image.yml` github action [#3444](https://github.com/chaos-mesh/chaos-mesh/pull/3444)
- Add a judgement of `enterNS` in `getAllInterfaces()` [#3459](https://github.com/chaos-mesh/chaos-mesh/pull/3459)
- Fix JVMChaos loading missing jar file for injection [#3491](https://github.com/chaos-mesh/chaos-mesh/pull/3491)

### Security

- Nothing

## [2.2.3] - 2022-08-04

### Added

- Add `QPS` and `Burst` for Chaos Dashboard Configuration [#3476](https://github.com/chaos-mesh/chaos-mesh/pull/3476)

### Changed

- Implement some missing methods in ctrl server [#3470](https://github.com/chaos-mesh/chaos-mesh/pull/3470)

### Deprecated

- Nothing

### Changed

- Nothing

### Removed

- Nothing

### Fixed

- Protect chaos available namespace and filter namespaces if needed [#3473](https://github.com/chaos-mesh/chaos-mesh/pull/3473)
- Respect flag `enableProfiling` and do not register profiler endpoints when it's false [#3474](https://github.com/chaos-mesh/chaos-mesh/pull/3474)
- Fix the blank screen after creating chaos experiment with "By YAML" [#3489](https://github.com/chaos-mesh/chaos-mesh/pull/3489)
- Update hint text about the manual token generating process for Kubernetes 1.24+ [#3505](https://github.com/chaos-mesh/chaos-mesh/pull/3505)

### Security

- Nothing

## [2.2.2] - 2022-07-07

### Changed

- Bump chaos-tproxy to v0.5.1 [#3426](https://github.com/chaos-mesh/chaos-mesh/pull/3426)
- Run build image ci on self-hosted machine [#3429](https://github.com/chaos-mesh/chaos-mesh/pull/3429)

## [2.2.1] - 2022-06-29

### Added

- JVMChaos: support inject fault into MySQL client [#3189](https://github.com/chaos-mesh/chaos-mesh/pull/3189)

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- PhysicalMachineChaos: update stress options type [#3354](https://github.com/chaos-mesh/chaos-mesh/pull/3354)
- StressChaos: run `pause` before `choom` [#3405](https://github.com/chaos-mesh/chaos-mesh/pull/3405)

### Security

- Nothing

## [2.2.0] - 2022-04-29

### Added

- Add metrics for archived objects in chaos-dashboard [#2568](https://github.com/chaos-mesh/chaos-mesh/pull/2568)
- Add metrics for iptables, ipset and tc metrics in chaos-daemon [#2540](https://github.com/chaos-mesh/chaos-mesh/pull/2540)
- Add metrics for emitted event counter in chaos-controller-manager [#2435](https://github.com/chaos-mesh/chaos-mesh/pull/2435)
- Add metrics for grpc client [#2458](https://github.com/chaos-mesh/chaos-mesh/pull/2458)
- Add metrics for grpc and HTTP request duration histogram [#2543](https://github.com/chaos-mesh/chaos-mesh/pull/2543)
- Add metrics for bpm controlled processes [#2497](https://github.com/chaos-mesh/chaos-mesh/pull/2497)
- Provide additional printer columns for `action` and `duration` [#2526](https://github.com/chaos-mesh/chaos-mesh/pull/2526)
- Add PhysicalMachine CRD [#2587](https://github.com/chaos-mesh/chaos-mesh/pull/2587)
- New command `physical-machine` to `chaosctl` [#2624](https://github.com/chaos-mesh/chaos-mesh/pull/2624)
- Add status "Deleting" for chaos experiments on Chaos Dashboard [#2708](https://github.com/chaos-mesh/chaos-mesh/pull/2708)
- Add time skew for gettimeofday [#2742](https://github.com/chaos-mesh/chaos-mesh/pull/2742)
- Add support of the Unified cgroup mode (tested with containerd runtime only) for linux stress experiments [#2928](https://github.com/chaos-mesh/chaos-mesh/pull/2928)
- Add `StatusCheck` CRD [#2954](https://github.com/chaos-mesh/chaos-mesh/pull/2954)
- Add support for declaring ports in external targets in NetworkChaos experiments [#2932](https://github.com/chaos-mesh/chaos-mesh/pull/2932)
- Add forced recovery of httpchaos, iochaos, stresschaos, and networkchaos for chaosctl [#2992](https://github.com/chaos-mesh/chaos-mesh/pull/2992)
- Add namespace and pod name in failed event for podxxxchaos crd [#3178](https://github.com/chaos-mesh/chaos-mesh/pull/3178)
- Add next generation `New Workflow` in UI [#3185](https://github.com/chaos-mesh/chaos-mesh/pull/3185)
- JVMChaos: support inject fault into MySQL client [#3189](https://github.com/chaos-mesh/chaos-mesh/pull/3189)

### Changed

- Use pipeline controller to serialize common controllers [#2465](https://github.com/chaos-mesh/chaos-mesh/pull/2465)
- Enable mTLS between chaos-controller-manager and chaosd [#2580](https://github.com/chaos-mesh/chaos-mesh/pull/2580)
- Rename Physics to Host in Chaos Dashboard [#2645](https://github.com/chaos-mesh/chaos-mesh/pull/2645)
- Retry oneshot chaos if it's not selected [#2618](https://github.com/chaos-mesh/chaos-mesh/pull/2618)
- Bump gopsutil to v3 [#2681](https://github.com/chaos-mesh/chaos-mesh/pull/2681)
- Add prefix for identifier of toda and tproxy in bpm [#2673](https://github.com/chaos-mesh/chaos-mesh/pull/2673)
- Bump toda to v0.2.2 [#2747](https://github.com/chaos-mesh/chaos-mesh/pull/2747)
- Bump go to 1.17 [#2754](https://github.com/chaos-mesh/chaos-mesh/pull/2754)
- Use github.com/pkg/errors to replace fmt.Errorf and "errors" [#2779](https://github.com/chaos-mesh/chaos-mesh/pull/2779)
- Kill chaos-tproxy while failing to apply config [#2672](https://github.com/chaos-mesh/chaos-mesh/pull/2672)
- JVMChaos: ignore AgentLoadException when install agent [#2701](https://github.com/chaos-mesh/chaos-mesh/pull/2701)
- Bump container-runtime to v0.11.0 [#2778](https://github.com/chaos-mesh/chaos-mesh/pull/2778)
- Bump kubernetes dependencies to v1.23.1 [#2778](https://github.com/chaos-mesh/chaos-mesh/pull/2778)
- Removed docker registry mirror [#2797](https://github.com/chaos-mesh/chaos-mesh/pull/2797)
- Use OpenAPI definitions to generate API Client and Form data in UI [2770](https://github.com/chaos-mesh/chaos-mesh/pull/2770)
- Refine logging in pkg/selector/pod [#3002](https://github.com/chaos-mesh/chaos-mesh/pull/3002)
- Add `envFollowKubernetesPattern` to handle k8s-like format env in helm templates [2955](https://github.com/chaos-mesh/chaos-mesh/pull/2955)
- Bump chaos-tproxy to v0.4.5 [#2555](https://github.com/chaos-mesh/chaos-mesh/pull/2555)
- Re-implement chaosctl based on ctrlserver [#2950](https://github.com/chaos-mesh/chaos-mesh/pull/2950)
- Fix wrong zero value of httpchaos replace-body-action[#2990](https://github.com/chaos-mesh/chaos-mesh/pull/2990)
- Bump gqlgen to v0.17.2 [#3038](https://github.com/chaos-mesh/chaos-mesh/pull/3038)
- Bump go to v1.18 [#3055](https://github.com/chaos-mesh/chaos-mesh/pull/3055)
- Bump toda to v0.2.3 [#3131](https://github.com/chaos-mesh/chaos-mesh/pull/3131)
- refactor: rename reconcileContext to reconcileInfo [#3154](https://github.com/chaos-mesh/chaos-mesh/pull/3154)
- Migrate e2e tests from self-hosted Jenkins to Github Action [#2986](https://github.com/chaos-mesh/chaos-mesh/pull/2986)
- Bump minimist from 1.2.5 to 1.2.6 in /ui [#3058](https://github.com/chaos-mesh/chaos-mesh/pull/3058)
- Specify image tag of `build-env` and `dev-env` for each branch [#3071](https://github.com/chaos-mesh/chaos-mesh/pull/3071)
- Specify image tag in e2e tests [#3147](https://github.com/chaos-mesh/chaos-mesh/pull/3147)
- Must update CHANGELOG [#3148](https://github.com/chaos-mesh/chaos-mesh/pull/3148)
- Use chaosDaemon.mtls.enabled instead of dashboard.securityMode for chaos-daemon mtls [#3168](https://github.com/chaos-mesh/chaos-mesh/pull/3168)
- Helm charts: component chaos-dashboard use certain service account and roles [#3145](https://github.com/chaos-mesh/chaos-mesh/pull/3145)
- Refactor helm charts template, split out webhook configuration and secrets [#3159](https://github.com/chaos-mesh/chaos-mesh/pull/3159)
- Helm charts: apply webhook.FailurePolicy to all the webhooks with default value `Fail` [#3184](https://github.com/chaos-mesh/chaos-mesh/pull/3184)
- Bump memStress from v0.2.1 to v0.3 [#3186](https://github.com/chaos-mesh/chaos-mesh/pull/3186)
- Helm charts: configure ca bundle for webhook explicitly [#3190](https://github.com/chaos-mesh/chaos-mesh/pull/3190)
- Refine logging in pkg/selector/generic/namespace [#3214](https://github.com/chaos-mesh/chaos-mesh/pull/3214)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Unable to load from saved objects [#2585](https://github.com/chaos-mesh/chaos-mesh/pull/2585)
- Fix helm install error [#2591](https://github.com/chaos-mesh/chaos-mesh/pull/2591)
- Fix helm conditions in ingress [#2604](https://github.com/chaos-mesh/chaos-mesh/pull/2604)
- Fix typo in NewExperiment [#2535](https://github.com/chaos-mesh/chaos-mesh/pull/2535)
- Fix chaos-kernel build, mark bcc version [#2693](https://github.com/chaos-mesh/chaos-mesh/pull/2693)
- Fix wrong field name of PhysicalMachineChaos on Chaos Dashboard [#2724](https://github.com/chaos-mesh/chaos-mesh/pull/2724)
- Fix field descriptions of GCPChaos [#2791](https://github.com/chaos-mesh/chaos-mesh/pull/2791)
- Fix `real_gettimeofday` on arm64 [#2849](https://github.com/chaos-mesh/chaos-mesh/pull/2849)
- Fix Github Action `upload-image` [#2935](https://github.com/chaos-mesh/chaos-mesh/pull/2935)
- Fix JVMChaos to handle the situation that the container which holds the JVM rules has been deleted [#2981](https://github.com/chaos-mesh/chaos-mesh/pull/2981)
- Fix typo in comments for Chaos API [#3109](https://github.com/chaos-mesh/chaos-mesh/pull/3109)

### Security

- Nothing

## [v2.1.8] - 2022-08-30

### Added

- Nothing

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix recover bug when setting force recover to true [#3578](https://github.com/chaos-mesh/chaos-mesh/pull/3578)
- Uniformly use Enter to add item in LabelField [#3580](https://github.com/chaos-mesh/chaos-mesh/pull/3580)

## [2.1.7] - 2022-07-29

### Added

- Add `QPS` and `Burst` for Chaos Dashboard Configuration [#3476](https://github.com/chaos-mesh/chaos-mesh/pull/3476)

### Changed

- Nothing

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Nothing

## [2.1.6] - 2022-07-22

### Added

- Add e2e build image cache [#3097](https://github.com/chaos-mesh/chaos-mesh/pull/3151)
- PhysicalMachineChaos: add recover command in process [#3468](https://github.com/chaos-mesh/chaos-mesh/pull/3468)

### Changed

- Must update CHANGELOG [#3181](https://github.com/chaos-mesh/chaos-mesh/pull/3181)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix: find the correct pid by CommName [#3336](https://github.com/chaos-mesh/chaos-mesh/pull/3336)

### Security

- Nothing

## [2.1.5] - 2022-04-18

### Added

- Nothing

### Changed

- Migrate e2e tests from self-hosted Jenkins to Github Action [#2986](https://github.com/chaos-mesh/chaos-mesh/pull/2986)
- Bump minimist from 1.2.5 to 1.2.6 in /ui [#3058](https://github.com/chaos-mesh/chaos-mesh/pull/3058)
- Specify image tag of `build-env` and `dev-env` for each branch [#3071](https://github.com/chaos-mesh/chaos-mesh/pull/3071)
- Bump toda to v0.2.3 [#3131](https://github.com/chaos-mesh/chaos-mesh/pull/3131)
- Specify image tag in e2e tests [#3147](https://github.com/chaos-mesh/chaos-mesh/pull/3147)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix `real_gettimeofday` on arm64 [#2849](https://github.com/chaos-mesh/chaos-mesh/pull/2849)
- Fix Github Action `upload-image` [#2935](https://github.com/chaos-mesh/chaos-mesh/pull/2935)
- Fix JVMChaos to handle the situation that the container which holds the JVM rules has been deleted [#2981](https://github.com/chaos-mesh/chaos-mesh/pull/2981)
- Fix typo in comments for Chaos API [#3109](https://github.com/chaos-mesh/chaos-mesh/pull/3109)

### Security

- Nothing

## [2.1.4] - 2022-03-21

### Added

- Add time skew for gettimeofday [#2742](https://github.com/chaos-mesh/chaos-mesh/pull/2742)

### Changed

- Removed docker registry mirror [#2797](https://github.com/chaos-mesh/chaos-mesh/pull/2797)

### Deprecated

- Nothing

### Removed

- Nothing

### Fixed

- Fix default value for concurrencyPolicy [#2622](https://github.com/chaos-mesh/chaos-mesh/pull/2622)
- Enable the webhooks for `Schedule` and `Workflow` [#2622](https://github.com/chaos-mesh/chaos-mesh/pull/2622)
- Fix PhysicalMachineChaos to make it able to create network bandwidth experiment. [#2850](https://github.com/chaos-mesh/chaos-mesh/pull/2850)
- Fix workflow emit new events after accomplished [#2911](https://github.com/chaos-mesh/chaos-mesh/pull/2911)
- Fix human unreadable logging timestamp [#2808](https://github.com/chaos-mesh/chaos-mesh/pull/2808) [#2902](https://github.com/chaos-mesh/chaos-mesh/pull/2902) [#2973](https://github.com/chaos-mesh/chaos-mesh/pull/2973)
- Fix default value of percent field in iochaos [#3018](https://github.com/chaos-mesh/chaos-mesh/pull/3018)
- Fix the unexpected CPU stress for StressChaos with cpu resource limit [#3102](https://github.com/chaos-mesh/chaos-mesh/pull/3102)
- Fix the bug that create JVMChaos failed in workflow [#3156](https://github.com/chaos-mesh/chaos-mesh/pull/3156)

### Security

- Nothing

## [2.1.3] - 2022-01-27

### Added

- Add status "Deleting" for chaos experiments on Chaos Dashboard [#2708](https://github.com/chaos-mesh/chaos-mesh/pull/2708)

### Changed

- Add prefix for identifier of toda and tproxy in bpm [#2673](https://github.com/chaos-mesh/chaos-mesh/pull/2673)
- Bump toda to v0.2.2 [#2747](https://github.com/chaos-mesh/chaos-mesh/pull/2747)
- Bump go to 1.17 [#2754](https://github.com/chaos-mesh/chaos-mesh/pull/2754)
- JVMChaos ignore AgentLoadException when install agent [#2701](https://github.com/chaos-mesh/chaos-mesh/pull/2701)
- Bump container-runtime to v0.11.0 [#2807](https://github.com/chaos-mesh/chaos-mesh/pull/2807)
- Bump kubernetes dependencies to v1.23.1 [#2807](https://github.com/chaos-mesh/chaos-mesh/pull/2807)
- Kill chaos-tproxy while failing to apply config [#2672](https://github.com/chaos-mesh/chaos-mesh/pull/2672)

### Fixed

- Fix wrong field name of PhysicalMachineChaos on Chaos Dashboard [#2724](https://github.com/chaos-mesh/chaos-mesh/pull/2724)
- Fix field descriptions of GCPChaos [#2791](https://github.com/chaos-mesh/chaos-mesh/pull/2791)
- Fix chaos experiment "not found" on Chaos Dashboard [#2698](https://github.com/chaos-mesh/chaos-mesh/pull/2698)

## [2.1.2] - 2021-12-29

### Changed

- Provide additional print columns for chaos experiments [#2526](https://github.com/chaos-mesh/chaos-mesh/pull/2526)
- Refactor pkg/time [#2570](https://github.com/chaos-mesh/chaos-mesh/pull/2570)
- Rename “physic” to “host” on Chaos Dashboard [#2645](https://github.com/chaos-mesh/chaos-mesh/pull/2645)
- Restructure UI codebase [#2590](https://github.com/chaos-mesh/chaos-mesh/pull/2590)
- Upgrade UI dependencies [#2685](https://github.com/chaos-mesh/chaos-mesh/pull/2685)
- Set default selector mode from “one” to “all” [#2680](https://github.com/chaos-mesh/chaos-mesh/pull/2792)
- Workflow now ordered by creation time [#2680](https://github.com/chaos-mesh/chaos-mesh/pull/2680)
- Set up codecov for testing coverage reports [#2679](https://github.com/chaos-mesh/chaos-mesh/pull/2679)
- Speed up e2e tests [#2617](https://github.com/chaos-mesh/chaos-mesh/pull/2617) [#2702](https://github.com/chaos-mesh/chaos-mesh/pull/2702)

### Fixed

- Fixed: error when using Schedule and PodChaos for injecting PodChaos as a cron job [#2618](https://github.com/chaos-mesh/chaos-mesh/pull/2618)
- Fixed: chaos-kernel build failure [#2693](https://github.com/chaos-mesh/chaos-mesh/pull/2693)

## [2.0.7] - 2022-01-27

### Added

- Add status "Deleting" for chaos experiments on Chaos Dashboard [#2708](https://github.com/chaos-mesh/chaos-mesh/pull/2708)

### Changed

- Add prefix for identifier of toda and tproxy in bpm [#2673](https://github.com/chaos-mesh/chaos-mesh/pull/2673)
- Kill chaos-tproxy while failing to apply config [#2672](https://github.com/chaos-mesh/chaos-mesh/pull/2672)

### Fixed

- Fix chaos experiment "not found" on Chaos Dashboard [#2698](https://github.com/chaos-mesh/chaos-mesh/pull/2698)
- Fix field descriptions of GCPChaos [#2791](https://github.com/chaos-mesh/chaos-mesh/pull/2791)

## [2.0.6] - 2021-12-29

### Changed

- Provide additional print columns for chaos experiments [#2526](https://github.com/chaos-mesh/chaos-mesh/pull/2526)
- Remove redundant codes [#2704](https://github.com/chaos-mesh/chaos-mesh/pull/2704)
- Speed up e2e tests #2617 [#2718](https://github.com/chaos-mesh/chaos-mesh/pull/2718)

### Fixed

- Fixed: error when using Schedule and PodChaos for injecting PodChaos as a cron job [#2618](https://github.com/chaos-mesh/chaos-mesh/pull/2618)
- Fixed: fail to recover when Chaos CR was deleted before appending finalizers [#2624](https://github.com/chaos-mesh/chaos-mesh/pull/2624)
- Fixed: chaos-kernel build failure [#2693](https://github.com/chaos-mesh/chaos-mesh/pull/2693)
- Fixed: Chaos Dashboard panic when creating StressChaos [#2655](https://github.com/chaos-mesh/chaos-mesh/pull/2655)
