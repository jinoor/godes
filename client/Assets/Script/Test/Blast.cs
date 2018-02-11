using UnityEngine;
using System.Collections;

public class Blast : MonoBehaviour {

    private float time_;

	// Use this for initialization
	void Start () {
        time_ = 0;
	
	}
	
	// Update is called once per frame
	void Update () {
        time_ += Time.deltaTime;
        if (time_ > 2.5)
        {
            Destroy(gameObject);
        }
	
	}
}
