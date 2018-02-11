using UnityEngine;
using System;

public class TestNet : MonoBehaviour {

    float time_ = 0f;

	// Use this for initialization
	void Start () 
    {
        NetClient.Instance.Connect();
    }
	
	// Update is called once per frame
	void Update () 
    {
        time_ += Time.deltaTime;
        if(time_ >= 3)
        {
            time_ = 0f;
        }
	}

    void OnApplicationQuit()
    {
        Debug.Log("退出游戏");
        NetClient.Instance.Close();
    }
}
