package stagelibrary

import (
	"errors"
	"github.com/streamsets/dataextractor/api"
	"github.com/streamsets/dataextractor/stages/destinations/http"
	"github.com/streamsets/dataextractor/stages/destinations/mqtt"
	"github.com/streamsets/dataextractor/stages/destinations/trash"
	"github.com/streamsets/dataextractor/stages/destinations/websocket"
	"github.com/streamsets/dataextractor/stages/origins/dev_random"
	"github.com/streamsets/dataextractor/stages/origins/filetail"
	"github.com/streamsets/dataextractor/stages/destinations/coap"
)

func CreateStageInstance(library string, stageName string) (api.Stage, error) {
	var instanceKey = library + ":" + stageName
	switch instanceKey {

	case "streamsets-datacollector-dev-lib:com_streamsets_pipeline_stage_devtest_RandomSource":
		return &dev_random.DevRandom{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_origin_logtail_FileTailDSource":
		return &filetail.FileTailOrigin{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_http_HttpClientDTarget":
		return &http.HttpClientDestination{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_websocket_WebSocketDTarget":
		return &websocket.WebSocketClientDestination{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_mqtt_MqttClientDTarget":
		return &mqtt.MqttClientDestination{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_coap_CoapClientDTarget":
		return &coap.CoapClientDestination{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_devnull_NullDTarget":
		return &trash.TrashDestination{}, nil

	case "streamsets-datacollector-basic-lib:com_streamsets_pipeline_stage_destination_devnull_ToErrorNullDTarget":
		return &trash.TrashDestination{}, nil
	}

	return nil, errors.New("No Stage Instance found for : " + instanceKey)
}
