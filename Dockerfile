FROM r351574nc3/bazel-onbuild:latest 

CMD ["./.output/execroot/__main__/bazel-out/k8-fastbuild/bin/application"]