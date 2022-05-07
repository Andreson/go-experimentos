#!/bin/sh


checkFail() {
 
retVal=$1
if [ $retVal -gt 0 ]; then
    echo "Erro nao esperado ao executar build "
    exit 1
fi

}

CONTAINER_NAME=cor-int-sidecar-cache-redis

if [ "$CONTAINER_REGISTRY_HOST" = "" ]; \
    then export CONTAINER_REGISTRY_HOST=andresonthiago; \
    fi
echo 'Gerando binario app em ./bin'

export CGO_ENABLED=0 
export GOOS=linux
export GOARCH=amd64

#parametros pra reduzir tamnho do binario
# -s remove informações de debug do executável 
# -w impede a geração do DWARF (Debugging With Attributed Record Formats).
#'-extldflags "-static"' compilação estatica baseada no SO destino
 
 go build -ldflags '-extldflags "-static" -s -w'  -o ./bin/sidecar-cache


# checkFail $?

 BUILD_VERSION=`md5sum ./bin/sidecar-cache | awk '{ print $1 }'`

 echo "Versao gerada $BUILD_VERSION"

 echo "iniciando geração  imagem docker "
 
 #docker build   . -t $CONTAINER_NAME

# checkFail $?

# echo "Imagem docker gerada :  "

# docker image ls $CONTAINER_NAME

# echo "Gerando tag imagem "

# docker tag $CONTAINER_NAME $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:$BUILD_VERSION

# docker tag $CONTAINER_NAME $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:latest

#docker tag $CONTAINER_NAME $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:2

# echo "Fazendo push imagens"

 #docker push  $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:$BUILD_VERSION

#docker push  $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:2

# checkFail $?

 #docker push  $CONTAINER_REGISTRY_HOST/$CONTAINER_NAME:latest

# checkFail $?

echo "########### build concluido #################"


