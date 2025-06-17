#!/bin/bash

ROOT_DIR="$(dirname "$0")/examples/samples/resources"
REPORT=""

# Colores para mejor visualización
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # Sin color

SUCCESS_FILE="terraform_examples_success.txt"
FAIL_FILE="terraform_examples_fail.txt"
SUMMARY_FILE="terraform_examples_summary.txt"

successes=()
failures=()

# Limpia archivos previos
> "$SUCCESS_FILE"
> "$FAIL_FILE"
> "$SUMMARY_FILE"

TIMEOUT=60

for dir in "$ROOT_DIR"/*/; do
    [ -d "$dir" ] || continue
    echo -e "${YELLOW}Entrando a: $dir${NC}"
    pushd "$dir" > /dev/null

    # Limpia archivos previos
    rm -rf .terraform .terraform.lock.hcl
    find . -maxdepth 1 -type f \( -name "*.log" -o -name "*.hcl" -o -name "*.info" -o -name "*.tfstate" -o -name "*.tfstate.backup" \) -exec rm -f {} +

    timeout $TIMEOUT terraform init -input=false > init.log 2>&1
    exit_code=$?
    if [ $exit_code -eq 124 ]; then
        echo -e "${RED}Timeout (>${TIMEOUT}s) en terraform init en $dir${NC}"
        failures+=("$dir: terraform init TIMEOUT (${TIMEOUT}s) (ver init.log)\n---")
        echo -e "$dir: terraform init TIMEOUT (${TIMEOUT}s) (ver init.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    elif [ $exit_code -ne 0 ]; then
        echo -e "${RED}Fallo terraform init en $dir${NC}"
        failures+=("$dir: terraform init failed (ver init.log)\n---")
        echo -e "$dir: terraform init failed (ver init.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    fi

    timeout $TIMEOUT terraform plan -input=false > plan.log 2>&1
    exit_code=$?
    if [ $exit_code -eq 124 ]; then
        echo -e "${RED}Timeout (>${TIMEOUT}s) en terraform plan en $dir${NC}"
        failures+=("$dir: terraform plan TIMEOUT (${TIMEOUT}s) (ver plan.log)\n---")
        echo -e "$dir: terraform plan TIMEOUT (${TIMEOUT}s) (ver plan.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    elif [ $exit_code -ne 0 ]; then
        echo -e "${RED}Fallo terraform plan en $dir${NC}"
        failures+=("$dir: terraform plan failed (ver plan.log)\n---")
        echo -e "$dir: terraform plan failed (ver plan.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    fi

    timeout $TIMEOUT terraform apply -auto-approve -input=false > apply.log 2>&1
    exit_code=$?
    if [ $exit_code -eq 124 ]; then
        echo -e "${RED}Timeout (>${TIMEOUT}s) en terraform apply en $dir${NC}"
        failures+=("$dir: terraform apply TIMEOUT (${TIMEOUT}s) (ver apply.log)\n---")
        echo -e "$dir: terraform apply TIMEOUT (${TIMEOUT}s) (ver apply.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    elif [ $exit_code -ne 0 ]; then
        echo -e "${RED}Fallo terraform apply en $dir${NC}"
        failures+=("$dir: terraform apply failed (ver apply.log)\n---")
        echo -e "$dir: terraform apply failed (ver apply.log)\n---" >> "../$FAIL_FILE"
        popd > /dev/null
        continue
    fi

    echo -e "${GREEN}Éxito en $dir${NC}"
    successes+=("$dir")
    echo "$dir" >> "../$SUCCESS_FILE"
    popd > /dev/null

done

# Resumen en archivo
{
    echo "==================== RESUMEN DE EJECUCIÓN ===================="
    echo ""
    echo "EXAMPLES EXITOSOS:"
    cat "$SUCCESS_FILE"
    echo ""
    echo "EXAMPLES FALLIDOS:"
    cat "$FAIL_FILE"
} > "$SUMMARY_FILE"

# Resumen en pantalla
if [ ${#failures[@]} -eq 0 ]; then
    echo -e "\n${GREEN}Todos los ejemplos se aplicaron correctamente.${NC}"
else
    echo -e "\n${RED}Resumen de fallos:${NC}"
    for fail in "${failures[@]}"; do
        echo -e "$fail"
    done
fi

echo -e "\nResumen guardado en $SUMMARY_FILE" 