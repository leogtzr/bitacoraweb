#!/bin/bash
#: author: Leo Gtz <leogutierrezramirez@gmail.com>

# set -x
readonly work_dir=$(dirname "$(readlink --canonicalize-existing "${0}")")
readonly bin_program="${work_dir}/web"
readonly error_program_not_found=80
readonly error_building_program=81

if [[ "${1}" == "b" ]]; then
    echo -n "building binary artifact ... "
    go clean
    go build -o "${work_dir}/web" || exit ${error_building_program}
fi

if [[ ! -f "${bin_program}" ]]; then
    echo "${0}: ${bin_program} does not exist." >&2
    exit ${error_program_not_found}
fi

echo "done"

export BITACORA_USER=leo
export BITACORA_PASSWORD=lein23

echo "running program ..."
"${bin_program}"

exit
