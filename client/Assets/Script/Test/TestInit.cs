using UnityEngine;
using System.Collections;

public class TestInit : MonoBehaviour {

    private Vector3 scale_ = new Vector3(10, 10, 10);
    private int dev_ = 50;
    

	// Use this for initialization
	void Start () {
        for (int i = 0; i < 5; ++i)
        {
            for (int j = 0; j < 5; ++j)
            {
                GameObject obj = GameObject.CreatePrimitive(PrimitiveType.Sphere);
                obj.transform.localScale = scale_;
                obj.transform.position = new Vector3(i * 20 - dev_, 0, j * 20 - dev_);
                LoginObj loginObj = obj.AddComponent<LoginObj>();
                loginObj.Init("name_" + i + "_" + j, "passwd_" + i + "_" + j);
                TestManager.Instance.lsLoginObj_.Add(loginObj);
            }
        }
	}
	
	// Update is called once per frame
	void Update () {
	
	}
}
