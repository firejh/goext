package gxinfluxdb

import (
	"github.com/influxdata/influxdb/client/v2"
	"time"
	"fmt"
)

type InfluxDbClient struct {
	//influxdb client
	cli 	client.Client
}

func CreateInfluxDbClient(host string, user string, pwd string) (InfluxDbClient, error){
	var err error
	var ret InfluxDbClient

	fmt.Println("&&&1:%v", ret)

	ret.cli, err =
		client.NewHTTPClient(client.HTTPConfig{
			Addr:     host,
			Username: user,
			Password: pwd,
		})
	if err != nil {
		return ret, err
	}
	fmt.Println("&&&2:%v", ret)
	return ret, nil
}

func (c InfluxDbClient)Write(batchPoint InfluxBatchPoints) error {
	err := c.cli.
		Write(batchPoint.b)
	if err != nil {
		return err
	}
	return nil
}

type InfluxBatchPoints struct {
	b 	client.BatchPoints
}

func CreateBatchPoints(database string, precision string) (*InfluxBatchPoints, error) {
	var (
		err error
		ret InfluxBatchPoints
	)
	ret = InfluxBatchPoints{}
	ret.b, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database: database,
		Precision: precision,
	})
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}

func (c InfluxBatchPoints)AddPoint(table string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	pt, err := client.NewPoint(table, tags, fields, t)
	if err != nil {
		return err
	}
	c.b.AddPoint(pt)
	return nil
}

func (c InfluxBatchPoints)Clear() {
	c.b.Clear()
}