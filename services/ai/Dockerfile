FROM tensorflow/tensorflow:1.12.0-py3

COPY ./service app

RUN pip install grpcio-tools
RUN pip3 install -r app/requirements.txt
WORKDIR app/
CMD ["python", "ai.py"]