FROM scratch
COPY hello /
COPY good-morning /
ENV PATH="/:${PATH}"
ENTRYPOINT ["/hello"]
CMD ["/good-morning"]