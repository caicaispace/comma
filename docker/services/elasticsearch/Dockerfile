ARG ELASTICSEARCH_VERSION
FROM elasticsearch:${ELASTICSEARCH_VERSION}

ARG ELASTICSEARCH_VERSION
ARG ELASTICSEARCH_PLUGINS

ENV PLUGINS=",${ELASTICSEARCH_PLUGINS},"

RUN if [[ -z "${PLUGINS##*,analysis-ik,*}" ]]; then \
        printf "y\n" | elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v${ELASTICSEARCH_VERSION}/elasticsearch-analysis-ik-${ELASTICSEARCH_VERSION}.zip; \
    fi \
    && \
    if [[ -z "${PLUGINS##*,analysis-pinyin,*}" ]]; then \
      printf "y\n" | elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-pinyin/releases/download/v${ELASTICSEARCH_VERSION}/elasticsearch-analysis-pinyin-${ELASTICSEARCH_VERSION}.zip; \
    fi

